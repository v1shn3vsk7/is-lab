import React, { useEffect, useState } from 'react';
import {useLocation, useNavigate} from 'react-router-dom';
import axios from 'axios';
import './AdminPage.css'

function AdminPage() {
    const location = useLocation()
    const [users, setUsers] = useState([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isPasswordModalOpen, setIsPasswordModalOpen] = useState(false)
    const [newUserLogin, setNewUserLogin] = useState('');
    const [isNewUserPasswordConstraint, setNewUserPasswordConstraint] = useState(false);
    const [passwordChangeSuccess, setPasswordChangeSuccess] = useState(false);
    const navigate = useNavigate();

    const [oldPassword, setOldPassword] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [repeatNewPassword, setRepeatNewPassword] = useState('');
    const [error, setError] = useState('');

    const { user_id } = location.state;

    useEffect(() => {
        axios.get('http://localhost:8080/api/list_users')
            .then((response) => {
                setUsers(response.data);
            })
            .catch((error) => {
                console.error('Error fetching user data:', error);
            });
    }, []);

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

        // if (is_password_constraint && !/^[A-Za-z0-9]+[.|,|;|:|-|!|?]+$/.test(newPassword)) {
        //     setError('На Ваш пароль действует ограничение: пароль должен иметь числа и знаки препинания');
        //     return;
        // }

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

    const openModal = () => {
        setIsModalOpen(true);
    };

    const closeModal = () => {
        setIsModalOpen(false);
    };

    const openPasswordModal = () => {
        setIsPasswordModalOpen(true);
        setPasswordChangeSuccess(false);
    }

    const createUser = async (login, isPasswordConstraint) => {
        const response = await axios.post('http://localhost:8080/api/create_user', {
            login: login,
            is_password_constraint: isPasswordConstraint,
        });

        if (response.status === 200) {
            window.location.reload()
        }
    };


    const handleExit = () => {
        navigate(`/`)
    };

    const handleSetPasswordRestriction = (userId, isPasswordConstraint) => {
        const request = {
            user_id: userId,
            is_password_constraint: !isPasswordConstraint,
        }

        fetch(`http://localhost:8080/api/update_user`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(request)
        })
            .then(() => {
                window.location.reload();
            })
            .catch((error) => console.error('Error setting password restriction:', error));
    };

    const handleBlockUser = (userId, isBlocked) => {
        const request = {
            user_id: userId,
            is_blocked: !isBlocked,
        }
        fetch(`http://localhost:8080/api/update_user`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(request)
        })
            .then(() => {
                window.location.reload();
            })
            .catch((error) => console.error('Error blocking user:', error));
    };

    const handleCreateUser = async (login, isPasswordConstraint) => {
        const response = await axios.post('http://localhost:8080/api/create_user', {
            login: login,
            is_password_constraint: isPasswordConstraint,
        });

        if (response.status === 200) {
            window.location.reload()
        }
    };

    // return (
    //     <div className="admin-container">
    //         {users.map((user) => (
    //             <div key={user.user_id} className="user-box">
    //                 <h3>{user.username}</h3>
    //                 <p>ID Пользователя: {user.user_id}</p>
    //                 <p>Заблокирован: {user.is_blocked ? 'Да' : 'Нет'}</p>
    //                 <p>Ограничение на пароль: {user.is_password_constraint ? 'Да' : 'Нет'}</p>
    //                 <button onClick={() => handleSetPasswordRestriction(user.user_id, user.is_password_constraint)}>{user.is_password_constraint ? "Выключить ограничение" : "Включить ограничение"}</button>
    //                 <button onClick={() => handleBlockUser(user.user_id, user.is_blocked)}>{user.is_blocked ? "Разблокировать": "Заблокировать"}</button>
    //             </div>
    //         ))}
    //         <button onClick={openModal}>Добавить пользователя</button>
    //         <div className="exit">
    //             <button onClick={handleExit} className="exit-button">Выход</button>
    //         </div>
    //     </div>
    //
    // );

    return (
        <div className="admin-container">
            {users.map((user) => (
                <div key={user.user_id} className="user-box">
                    <h3>{user.username}</h3>
                    <p>Логин: {user.login}</p>
                    <p>Заблокирован: {user.is_blocked ? 'Да' : 'Нет'}</p>
                    <p>Ограничение на пароль: {user.is_password_constraint ? 'Да' : 'Нет'}</p>
                    <button onClick={() => handleSetPasswordRestriction(user.user_id, user.is_password_constraint)}>{user.is_password_constraint ? "Выключить ограничение" : "Включить ограничение"}</button>
                    <button onClick={() => handleBlockUser(user.user_id, user.is_blocked)}>{user.is_blocked ? "Разблокировать": "Заблокировать"}</button>
                </div>
            ))}
            <div className="exit">
                <button onClick={handleExit} className="exit-button">
                    Выход
                </button>
            </div>
            <div className="create-user-button">
                <button onClick={() => setIsModalOpen(true)}>Create User</button>
                <button onClick={openPasswordModal}>Change Password</button>
            </div>
            {isModalOpen && (
                <div className="modal">
                    <div className="modal-content">
                        <h2>Create User</h2>
                        <input
                            type="text"
                            placeholder="User Login"
                            value={newUserLogin}
                            onChange={(e) => setNewUserLogin(e.target.value)}
                        />
                        <label>
                            <input
                                type="checkbox"
                                checked={isNewUserPasswordConstraint}
                                onChange={() =>
                                    setNewUserPasswordConstraint(!isNewUserPasswordConstraint)
                                }
                            />
                            Password Constraint
                        </label>
                        <button onClick={() => handleCreateUser(newUserLogin, isNewUserPasswordConstraint)}>Create</button>
                        <button onClick={() => setIsModalOpen(false)}>Cancel</button>
                    </div>
                </div>
            )}

            {isPasswordModalOpen && (
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

export default AdminPage;
