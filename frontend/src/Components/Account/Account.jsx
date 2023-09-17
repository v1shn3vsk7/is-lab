import React, { useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import './Account.css'; // Import the CSS file
import axios from 'axios';

function Account() {
    const location = useLocation();
    const navigate = useNavigate()
    const { user_id, username, is_blocked, is_password_constraint } = location.state;

    // State for the modal and form fields
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [oldPassword, setOldPassword] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [repeatNewPassword, setRepeatNewPassword] = useState('');
    const [passwordChangeSuccess, setPasswordChangeSuccess] = useState(false);
    const [error, setError] = useState('');

    // Function to open the modal
    const openModal = () => {
        setIsModalOpen(true);
        setPasswordChangeSuccess(false); // Reset success message
        setError('');
    };

    // Function to close the modal
    const closeModal = () => {
        setIsModalOpen(false);
        setOldPassword('');
        setNewPassword('');
        setRepeatNewPassword('');
        setPasswordChangeSuccess(false); // Reset success message
        setError('');
    };

    // Function to handle password change
    const handlePasswordChange = async () => {
        try {
            const response = await axios.post('http://localhost:8080/api/change_password', {
                user_id,
                old_password: oldPassword,
                new_password: newPassword,
            });

            if (response.status === 200) {
                setPasswordChangeSuccess(true);
                setError('');
                closeModal();
            }
        } catch (error) {
            setError(error.message);
        }
    };

    // Function to handle logout
    const handleLogout = () => {
        navigate(`/`)
    };

    return (
        <div className="account-container"> {/* Apply the container class */}
            <h1>Account</h1>
            <p className="account-item">User ID: {user_id}</p>
            <p className="account-item">Username: {username}</p>
            <p className="account-item">
                Is Blocked: <span className={is_blocked ? 'yes' : 'no'}>{is_blocked ? 'Yes' : 'No'}</span>
            </p>
            <p className="account-item">
                Has Password Constraint: <span className={is_password_constraint ? 'yes' : 'no'}>{is_password_constraint ? 'Yes' : 'No'}</span>
            </p>

            {/* Button to open the password change modal */}
            <button onClick={openModal}>Change Password</button>

            {/* Button to log out */}
            <button onClick={handleLogout} className="logout-button">Log Out</button>

            {/* Modal for changing password */}
            {isModalOpen && (
                <div className="modal">
                    <div className="modal-content">
                        <h2>Change Password</h2>
                        <input
                            type="password"
                            placeholder="Old Password"
                            value={oldPassword}
                            onChange={(e) => setOldPassword(e.target.value)}
                        />
                        <input
                            type="password"
                            placeholder="New Password"
                            value={newPassword}
                            onChange={(e) => setNewPassword(e.target.value)}
                        />
                        <input
                            type="password"
                            placeholder="Repeat New Password"
                            value={repeatNewPassword}
                            onChange={(e) => setRepeatNewPassword(e.target.value)}
                        />
                        <button onClick={handlePasswordChange}>Save</button>
                        <button onClick={closeModal}>Cancel</button>
                        {passwordChangeSuccess && <p>Password changed successfully!</p>}
                        {error && <p>Error: {error}</p>}
                    </div>
                </div>
            )}
        </div>
    );
}

export default Account;
