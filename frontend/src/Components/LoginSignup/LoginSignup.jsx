import React, {useState} from 'react'
import { useNavigate } from 'react-router-dom';
import './LoginSignup.css'
import user_icon from '../Assets/person.png'
import email_icon from '../Assets/email.png'
import password_icon from '../Assets/password.png'
import axios from "axios";


const LoginSignup = () => {
    const [action,setAction] = useState("Регистрация")
    const [fio, setFio] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const handleSubmit = async (event) => {
        event.preventDefault();
        setError(null)

        if (event.target.innerText === 'Вход') {
            if (action === "Регистрация") {
                setAction("Вход")
                return
            }

            if (email) {
                try {
                    const response = await axios.get('http://localhost:8080/api/find_user', {
                        params: {
                            login: email,
                            password: password,
                        }
                    });

                    if (response.status === 404) {
                        setError("user is not found")
                        return
                    }

                    const userData = response.data

                    if (userData.is_blocked) {
                        window.location.href = '/blocked-page';
                    }

                    if (userData.is_empty_password) {
                        navigate(`/setup-password/${userData.user_id}`, {
                            state: {
                                user_id: userData.user_id,
                                is_password_constraint: userData.is_password_constraint,
                                username: userData.username,
                                is_blocked: userData.is_blocked,
                                login: email,
                            }
                        })
                        return
                    }

                    if (email === "admin") {
                        console.log("ADMIN DETECTED")
                        navigate(`/admin-page/${userData.user_id}`, {
                            state: {
                                user_id: userData.user_id,
                                username: userData.username,
                                is_blocked: userData.is_blocked,
                                is_password_constraint: userData.is_password_constraint,
                            }
                        })
                        return
                    }

                    if (userData.is_password_constraint && !/^[A-Za-z0-9]+[.|,|;|:|-|!|?]+$/.test(password)) {
                        navigate(`/setup-password/${userData.user_id}`, {
                            state: {
                                user_id: userData.user_id,
                                is_password_constraint: userData.is_password_constraint,
                                username: userData.username,
                                is_blocked: userData.is_blocked,
                                login: email,
                            }
                        })
                        return
                    }

                    if (userData) {
                        navigate(`/account/${userData.user_id}`, {
                            state: {
                                user_id: userData.user_id,
                                username: userData.username,
                                is_blocked: userData.is_blocked,
                                is_password_constraint: userData.is_password_constraint,
                                is_empty_password: userData.is_empty_password,
                                login: email,
                            },
                        });
                    }

                } catch (error) {
                    if (error.response && error.response.status === 404 || error.response.status === 400) {
                        setError('Неверный логин и/или пароль')
                        return
                    }

                    setError(error.message)
                }

            } else {
                setError('Пожалуйста, заполните все поля');
            }
        }
    };

    return (
        <div className='container'>
            <div className="header">
                <div className="text">{action}</div>
                <div className="underline"></div>
            </div>

            <div className="inputs">
                {action==="Вход"?<div></div>:<div className="input">
                    <img src={user_icon} alt=""/>
                    <input
                        type="text"
                        placeholder="ФИО"
                        value={fio}
                        onChange={(event) =>
                            setFio(event.target.value)}
                    />
                </div>}

                <div className="input">
                    <img src={email_icon} alt=""/>
                    <input
                        type="email"
                        placeholder="Логин"
                        value={email}
                        onChange={(event) =>
                            setEmail(event.target.value)}
                    />
                </div>

                <div className="input">
                    <img src={password_icon} alt=""/>
                    <input
                        type="password"
                        placeholder="Пароль"
                        value={password}
                        onChange={(event) =>
                        setPassword(event.target.value)}
                    />
                </div>

            </div>

            <div className="submit-container">
                <div
                    className={action === 'Вход' ? 'submit gray' : 'submit'}
                    onClick={handleSubmit}
                >
                    {action === 'Вход' ? 'Регистрация' : 'Регистрация'}
                </div>
                <div
                    className={action === 'Регистрация' ? 'submit gray' : 'submit'}
                    onClick={handleSubmit}
                >
                    {action === 'Регистрация' ? 'Вход' : 'Вход'}
                </div>
            </div>

            {error && <h3 className="error">{error}</h3>}

        </div>
    )
}

export default LoginSignup