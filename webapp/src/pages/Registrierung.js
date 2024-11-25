import React, { useState } from 'react';
import axios from 'axios';
import '../kistenmeister.css';
import { Table } from 'react-bootstrap';
import * as Icon from 'react-bootstrap-icons';
import Config from '../km-config';

function Registrierung() {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [passwordVisible, setPasswordVisible] = useState(false);
    const [lizenz, setLizenz] = useState('');
    const [message, setMessage] = useState('');

    const handlePasswordVisibility = () => {
        setPasswordVisible(!passwordVisible);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/register', { name, email, password, lizenz });
            setMessage(response.data.message);
            window.location.href = "/aktivierung";
        } catch (error) {
            setMessage(error.response.data.error);
        }
    };

    return (
        <div className="App">
            <header className="App-header">
                <div className="container">
                    <h1 className="my-4 text-center">
                        <Icon.PersonPlus className="me-2" />
                        Kistenmeister
                    </h1>
                    <div className="d-flex justify-content-center">
                        <form onSubmit={handleSubmit} className="mb-3">
                            <Table className="table table-bordered">
                                <tbody className="km-logintable">
                                    <tr>
                                        <td><label htmlFor="name"><Icon.Person className="me-2" />Name:</label></td>
                                        <td>
                                            <input
                                                type="text"
                                                id="name"
                                                className="form-control form-control-sm"
                                                value={name}
                                                onChange={(e) => setName(e.target.value)}
                                                required
                                                style={{ width: '400px' }}
                                            />
                                        </td>
                                    </tr>
                                    <tr>
                                        <td><label htmlFor="lizenz"><Icon.Key className="me-2" />Lizenz:</label></td>
                                        <td>
                                            <select
                                                id="lizenz"
                                                className="form-control form-control-sm"
                                                value={lizenz}
                                                onChange={(e) => setLizenz(e.target.value)}
                                                required
                                                style={{ width: '400px' }}>
                                                <option value=""> </option>
                                                <option value="Pro">Pro</option>
                                                <option value="User">User</option>
                                            </select>
                                        </td>
                                    </tr>
                                    <tr>
                                        <td><label htmlFor="email"><Icon.Envelope className="me-2" />Email:</label></td>
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
                                        <td><label htmlFor="password"><Icon.Lock className="me-2" />Passwort:</label></td>
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
                                <button type="submit" className="btn btn-primary mb-2">
                                    <Icon.PersonCheck className="me-2" />
                                    Register
                                </button>
                                <br />
                                <button onClick={() => window.location.href = "/"} className="btn btn-secondary">
                                    <Icon.BoxArrowInRight className="me-2" />
                                    Login
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

export default Registrierung;
