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

            if (email && password) {
                try {
                    const response = await axios.get('http://localhost:8080/api/find_user', {
                        params: {
                            login: email,
                            password: password,
                        }
                    });

                    const userData = response.data

                    if (userData) {
                        navigate(`/account/${userData.ID}`, {
                            state: {
                                user_id: userData.user_id,
                                username: userData.username,
                                is_blocked: userData.is_blocked,
                                is_password_constraint: userData.is_password_constraint,
                            },
                        });
                    }

                } catch (error) {
                    setError(error.message)
                }

            } else {
                setError('Пожалуйста, заполните все поля');
            }
        } else if (event.target.innerText === 'Регистрация') {
            if (action === "Вход") {
                setAction("Регистрация")
                return
            }

            if (email && password && fio) {
                const request = {
                    username: fio,
                    login: email,
                    password: password
                }

                fetch('http://localhost:8080/api/signup', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(request)
                }).then(response => {
                    if (response.ok) {
                        console.log("login OK")
                    } else {
                        setError(response.get("error"))
                    }
                }).catch(error => {
                    console.error('Error:', error);
                })

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