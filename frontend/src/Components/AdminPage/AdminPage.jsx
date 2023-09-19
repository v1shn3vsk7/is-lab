import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import './AdminPage.css'

function AdminPage() {
    const [users, setUsers] = useState([]);
    const navigate = useNavigate()

    useEffect(() => {
        axios.get('http://localhost:8080/api/list_users')
            .then((response) => {
                setUsers(response.data);
            })
            .catch((error) => {
                console.error('Error fetching user data:', error);
            });
    }, []);

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

    return (
        <div className="admin-container">
            {users.map((user) => (
                <div key={user.user_id} className="user-box">
                    <h3>{user.username}</h3>
                    <p>ID Пользователя: {user.user_id}</p>
                    <p>Заблокирован: {user.is_blocked ? 'Да' : 'Нет'}</p>
                    <p>Ограничение на пароль: {user.is_password_constraint ? 'Да' : 'Нет'}</p>
                    <button onClick={() => handleSetPasswordRestriction(user.user_id, user.is_password_constraint)}>{user.is_password_constraint ? "Выключить ограничение" : "Включить ограничение"}</button>
                    <button onClick={() => handleBlockUser(user.user_id, user.is_blocked)}>{user.is_blocked ? "Разблокировать": "Заблокировать"}</button>
                </div>
            ))}
            <div className="exit">
                <button onClick={handleExit} className="exit-button">Выход</button>
            </div>
        </div>

    );
}

export default AdminPage;
