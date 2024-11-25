import React, { useEffect, useState } from 'react';
import axios from 'axios';
import '../kistenmeister.css';
import { Table, Button } from 'react-bootstrap';
import * as Icon from 'react-bootstrap-icons';
import { useContext } from 'react';
import { UserContext } from '../App';
import Config from '../km-config';

function Home() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [passwordVisible, setPasswordVisible] = useState(false);
    const [message, setMessage] = useState('');

    const handlePasswordVisibility = () => {
        setPasswordVisible(!passwordVisible);
    };

    const { user, setUser } = useContext(UserContext);

    function getUserData(ptoken) {

        fetch (Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/person',
        {
            method: "GET",
            headers: { 'Authorization': ("Bearer " + ptoken) }
        }
        )
            .then(response => response.json())
            .then(result => {

                localStorage.setItem("id", result.Person.ID);
                localStorage.setItem("name", result.Person.Name);
                localStorage.setItem("email", result.Person.Email);
                localStorage.setItem("token", result.Person.Token);

                window.location.href = "/list";
            });
    }

    function handleSubmit(e) {
        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password}),
        })
            .then(response => response.json())
            .then(data => {
                console.log('Success fetching:', data);
                setMessage('Login erfolgreich');
                getUserData(data.Token);
            })
            .catch((error) => {
                setMessage('Fehler beim Login');
                console.error('Error:', error);
            });
    }

    useEffect(() => {
        console.log("User: " + JSON.stringify(user));
      }, [user]);

    return (
        <div className="App">
            <header className="App-header">
                <div className="container">
                    <h1 className="my-4 text-center">
                        <Icon.BoxSeam className="me-2" />
                        Kistenmeister
                    </h1>
                    <div className="d-flex justify-content-center">
                        <form onSubmit={handleSubmit} className="mb-3">
                            <Table className="table table-bordered">
                                <tbody className="km-logintable">
                                    <tr>
                                        <td><label htmlFor="email"><Icon.Envelope className="me-1" />Email:</label></td>
                                        <td>
                                            <input
                                                type="email"
                                                id="email"
                                                className="form-control form-control-sm"
                                                value={email}
                                                onChange={(e) => setEmail(e.target.value)}
                                                required
                                                style={{ width: '400px' }}
                                            />
                                        </td>
                                    </tr>
                                    <tr>
                                        <td><label htmlFor="password"><Icon.Lock className="me-1" />Passwort:</label></td>
                                        <td>
                                            <div className="input-group">
                                                <input
                                                    type={passwordVisible ? 'text' : 'password'}
                                                    id="password"
                                                    className="form-control form-control-sm"
                                                    value={password}
                                                    onChange={(e) => setPassword(e.target.value)}
                                                    required
                                                />
                                                <div className="input-group-append">
                                                    <button
                                                        type="button"
                                                        className="btn btn-outline-secondary btn-sm"
                                                        onClick={handlePasswordVisibility}
                                                    >
                                                        {passwordVisible ? <Icon.EyeSlash /> : <Icon.Eye />}
                                                    </button>
                                                </div>
                                            </div>
                                        </td>
                                    </tr>
                                </tbody>
                            </Table>
                            <div className="text-center">
                                <Button onClick={() => handleSubmit()} className="btn btn-primary mb-1">
                                    <Icon.BoxArrowInRight className="me-2" />
                                    Login
                                </Button>
                                <br />
                                <button onClick={() => window.location.href = "/register"} className="btn btn-secondary">
                                    <Icon.PersonPlus className="me-2" />
                                    Register
                                </button>
                            </div>
                        </form>
                    </div>
                    {message && <div className="alert alert-info text-center">{message}</div>}
                </div>
            </header>
        </div>
    );
}

export default Home;
