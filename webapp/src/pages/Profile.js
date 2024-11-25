import React, { useContext, useEffect, useState } from 'react';
import { Container, Form, Table, Button, Alert } from 'react-bootstrap';
import Header from '../components/Header';
import * as Icon from 'react-bootstrap-icons';
import { UserContext } from '../App';
import Config from '../km-config';

function Profile() {

    const [editMode, setEditMode] = useState(false);
    const { user, setUser } = useContext(UserContext);

    const [showAlert, setShowAlert] = useState(false);
    const [alertVariant, setAlertVariant] = useState('');
    const [alertText, setAlertText] = useState('');

    function saveProfile() {
        if (document.getElementById('Passwort').value !== document.getElementById('rePasswort').value) {
            setAlertVariant('danger');
            setAlertText('Die Passwörter stimmen nicht überein.');
            setShowAlert(true);
            return;
        }
        setShowAlert(false);

    }

    function deleteProfile() {
        console.log("Delete Box");
        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/person/' + user.id,
        {
            method: "DELETE",
            headers: { 'Authorization': ("Bearer " + user.token) }
        }
        )
            .then(response => response.text())
            .then(result => {
                console.log(result);
                window.location.href = '/';
            });
    }

    useEffect(() => {
        
    }, []);

    return (
        <div>
            <Header />
            <Container className='km-page-content'>
                { showAlert &&
                    <Alert variant={alertVariant}>
                        {alertText}
                    </Alert>
                }
                <Form>
                    <h1>Profil von {user.name}</h1>
                    <Table>
                        <tbody>
                            <tr>
                                <td>
                                    <Form.Label>Name: </Form.Label>
                                </td>
                                <td>
                                    { !editMode &&
                                        <div>{user.name}</div>
                                    }
                                    { editMode &&
                                        <Form.Group controlId="Name">
                                            <Form.Control placeholder={user.name} />
                                        </Form.Group>
                                    }
                                </td>
                                
                            </tr>
                            <tr>
                                <td>
                                    <Form.Label>Email: </Form.Label>
                                </td>
                                <td>
                                    { !editMode &&
                                        <div>{user.email}</div>
                                    }
                                    { editMode &&
                                        <Form.Group controlId="Email">
                                            <Form.Control type='email' placeholder={user.email} />
                                        </Form.Group>
                                    }
                                </td>
                            </tr>
                            { editMode &&
                                <tr>
                                    <td>
                                        <Form.Label>Passwort: </Form.Label>
                                    </td>
                                    <td>
                                        <Form.Group controlId="Passwort">
                                            <Form.Control type='password' placeholder="Passwort" />
                                        </Form.Group>
                                    </td>
                                </tr>
                                
                            }
                            { editMode &&
                                <tr>
                                    <td>
                                        <Form.Label>Passwort wiederholen: </Form.Label>
                                    </td>
                                    <td>
                                        <Form.Group controlId="rePasswort">
                                            <Form.Control type='password' placeholder="Passwort wiederholen" />
                                        </Form.Group>
                                    </td>
                                </tr>
                                
                            }
                        </tbody>
                    </Table>
                    { !editMode &&
                        <div>
                            <Button className="km-btn-in-list" variant="secondary" onClick={() => setEditMode(true)}><Icon.Pencil /> Bearbeiten</Button>
                        </div>
                    }
                    { editMode &&
                        <div>
                            <Button className="km-btn-in-list" variant="primary" onClick={() => saveProfile()}><Icon.Save /> Speichern</Button>
                            <Button className="km-btn-in-list" variant="danger" onClick={() => deleteProfile()}><Icon.Trash /> Profil löschen</Button>
                            <Button className="km-btn-in-list" variant="secondary" onClick={() => setEditMode(false)}><Icon.SignStop /> Abbrechen</Button>
                        </div>
                    }
                </Form>
            </Container>
        </div>
    );
}
export default Profile;