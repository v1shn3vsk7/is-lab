import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './BlockedPage.css';

function BlockedPage() {
    const navigate = useNavigate()

    const handleExit = () => {
        navigate(`/`)
    };

    return (
        <div className="blocked-info">
            <h3>Ваш аккаунт заблокирован администратором</h3>
            <button onClick={handleExit} className="exit-button">Вернуться</button>
        </div>
    )
}

export default BlockedPage