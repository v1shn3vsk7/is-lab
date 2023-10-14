import React, { useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import axios from 'axios';

function SetupPassword() {
    const location = useLocation();
    const navigate = useNavigate();
    const { user_id, is_password_constraint, username, is_blocked, login } = location.state;

    const [newPassword, setNewPassword] = useState('');
    const [repeatNewPassword, setRepeatNewPassword] = useState('');
    const [passwordChangeSuccess, setPasswordChangeSuccess] = useState(false);
    const [error, setError] = useState('');

    const handlePasswordChange = async () => {
        if (!(newPassword && repeatNewPassword)) {
            setError('Заполните все поля');
            return;
        }

        if (newPassword !== repeatNewPassword) {
            setError('Пароли не совпадают');
            return;
        }

        if (is_password_constraint && !/^[A-Za-z0-9]+[.|,|;|:|-|!|?]+$/.test(newPassword)) {
            setError('На Ваш пароль действует ограничение: пароль должен иметь числа и знаки препинания');
            return;
        }

        // Add any additional validation logic here, such as password constraints.

        try {
            const response = await axios.post('http://localhost:8080/api/setup_password', {
                user_id: user_id,
                password: newPassword,
            });

            if (response.status === 200) {
                setPasswordChangeSuccess(true);
                setError('');

                window.location.href = '/';
                return

                if (login === "admin") {
                    navigate(`/admin-page/${user_id}`, {
                        state: {
                            user_id: user_id,
                            username: username,
                            is_blocked: is_blocked,
                            is_password_constraint: is_password_constraint,
                        }
                    })
                    return
                }

                navigate(`/account/${user_id}`, {
                    state: {
                        user_id: user_id,
                        username: username,
                        is_blocked: is_blocked,
                        is_password_constraint: is_password_constraint,
                    },
                });
            }
        } catch (error) {
            setError(error.message);
        }
    };

    return (
        <div className="change-password-container">
            <h2>Change Password</h2>
            <input
                type="password"
                placeholder="Новый пароль"
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
            />
            <input
                type="password"
                placeholder="Повторите новый пароль"
                value={repeatNewPassword}
                onChange={(e) => setRepeatNewPassword(e.target.value)}
            />
            <button onClick={handlePasswordChange}>Сохранить</button>
            {passwordChangeSuccess && <p>Пароль изменен!</p>}
            {error && <p>Error: {error}</p>}
        </div>
    );
}

export default SetupPassword