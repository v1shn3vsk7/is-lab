import React, { useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import './Account.css';
import axios from 'axios';

function Account() {
    const location = useLocation();
    const navigate = useNavigate()
    const { user_id, username, is_blocked, is_password_constraint } = location.state;

    // prevent user to enter blocked account directly from url
    if (is_blocked) {
        window.location.href = '/blocked-page';
    }

    const [isModalOpen, setIsModalOpen] = useState(false);
    const [oldPassword, setOldPassword] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [repeatNewPassword, setRepeatNewPassword] = useState('');
    const [passwordChangeSuccess, setPasswordChangeSuccess] = useState(false);
    const [error, setError] = useState('');

    const openModal = () => {
        setIsModalOpen(true);
        setPasswordChangeSuccess(false);
        setError('');
    };

    const closeModal = () => {
        setIsModalOpen(false);
        setOldPassword('');
        setNewPassword('');
        setRepeatNewPassword('');
        setPasswordChangeSuccess(false);
        setError('');
    };

    const handlePasswordChange = async () => {
        if (!(oldPassword && repeatNewPassword && newPassword)) {
            setError('Заполните все поля')
            return
        }

        if (newPassword !== repeatNewPassword) {
            setError('Пароли не совпадают')
            return
        }

        if (newPassword === oldPassword) {
            setError('Такой пароль уже установлен')
            return
        }

        if (is_password_constraint && !/^[A-Za-z0-9]+[.|,|;|:|-|!|?]+$/.test(newPassword)) {
            setError('На Ваш пароль действует ограничение: пароль должен иметь числа и знаки препинания');
            return;
        }

        try {
            const response = await axios.post('http://localhost:8080/api/update_user_password', {
                user_id: user_id,
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

    const handleLogout = () => {
        navigate(`/`)
    };

    return (
        <div className="account-container"> {}
            <h1>Account</h1>
            <p className="account-item">ID Пользователя: {user_id}</p>
            <p className="account-item">ФИО: {username}</p>
            <p className="account-item">
                Заблокирован: <span className={is_blocked ? 'Да' : 'Нет'}>{is_blocked ? 'Да' : 'Нет'}</span>
            </p>
            <p className="account-item">
                Ограничение на пароль: <span className={is_password_constraint ? 'Да' : 'Нет'}>{is_password_constraint ? 'Да' : 'Нет'}</span>
            </p>

            <button onClick={openModal}>Change Password</button>

            <button onClick={handleLogout} className="logout-button">Log Out</button>

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
