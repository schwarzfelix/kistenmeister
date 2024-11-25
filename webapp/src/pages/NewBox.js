import React from 'react';
import Header from '../components/Header';
import { Container, Form, Button, Alert, Spinner } from 'react-bootstrap';
import * as Icon from 'react-bootstrap-icons';
import { useState } from 'react';
import Config from '../km-config';
import { useContext } from 'react';
import { UserContext } from '../App';

function NewBox() {

    const [showLocationSipnner, setShowLocationSpinner] = useState(false);

    const [showAlert, setShowAlert] = useState(false);
    const [alertVariant, setAlertVariant] = useState('');
    const [alertText, setAlertText] = useState('');

    const { user, setUser } = useContext(UserContext);

    let addFav = (id) => {
        const data = new FormData();
        data.append("kiste_id", id);

        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/merklisteneinträge/YKz7gWpxSMbVYh8XNbMBRJMMTiBgtg3HPPvRlUEY_cQ=/',
            {
                method: "POST",
                headers: { 'Authorization': ("Bearer " + user.token) },
                body: data,
                mode: 'no-cors'
            }
            )
            .then(response => response.text())
            .then(result => {
                console.log('Success:', result);
            }
        );
    }

    function save() {
        console.log('save');

        const name = document.getElementById('name').value;
        const beschreibung = document.getElementById('beschreibung').value;
        const ort = document.getElementById('ort').value;

        const data = new FormData();
        data.append('Name', name);
        data.append('Beschreibung', beschreibung);
        data.append('Ort', ort);

        console.log(data);

        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kiste', {
            method: 'POST',
            headers: { 'Authorization': ("Bearer " + user.token) },
            body: data
        })
            .then(response => response.json())
            .then(data => {
                console.log('Success:', data);
                addFav(data["Neue Kiste mit ID"]);
                window.location.href = ('/box/' + data["Neue Kiste mit ID"]);
            })
            .catch((error) => {
                console.error('Error:', error);
            });

    }


    function updateLocation() {
        setShowLocationSpinner(true);
        navigator.geolocation.getCurrentPosition((position) => {

            //check permission
            if (!navigator.permissions) {
                
                setAlertText("Ihr Browser unterstützt keine Standortbestimmung!");
                setAlertVariant("danger");
                setShowAlert(true);

            } else {
                navigator.permissions.query({ name: 'geolocation' }).then((result) => {
                    if (result.state === 'granted') {

                        const latitude = position.coords.latitude;
                        const longitude = position.coords.longitude;
                        const location = latitude + ", " + longitude;
                        document.getElementById('ort').value = location;

                        setAlertText("Standort erfolgreich bestimmt!");
                        setAlertVariant("success");
                        setShowAlert(true);
                        setShowLocationSpinner(false);

                    } else if (result.state === 'prompt') {

                        setAlertText("Bitte erlauben Sie den Zugriff auf Ihren Standort!)");
                        setAlertVariant("warning");
                        setShowAlert(true);

                    } else if (result.state === 'denied') {

                        setAlertText("Zugriff auf Standort verweigert!");
                        setAlertVariant("danger");
                        setShowAlert(true);

                    }
                    result.onchange = function () {
                        console.log(result.state);
                    };
                });
            }
        }
        );
    }

    return (
        <div>
            <Header />
            <Container className='km-page-content'>
                { showAlert &&
                    <div className='km-section'>
                        <Alert variant={alertVariant}>
                            {alertText}
                        </Alert>
                    </div>
                }
                <h1>Neue Kiste erstellen</h1>
                <Form>
                    <div className='km-form-space'>
                        <Form.Group controlId="name">
                            <Form.Control type="text" size="lg" placeholder="Name" />
                        </Form.Group>
                    </div>
                    <div className='km-form-space'>
                        <Form.Group controlId="beschreibung">
                            <Form.Control as="textarea" rows={3} placeholder="Beschreibung" />
                        </Form.Group>
                    </div>
                    <div className='km-form-space'>
                        <Form.Group controlId="ort">
                            <Form.Control type="text" size="sm" placeholder="Ort" />
                        </Form.Group>
                        <Button size='sm' variant="light" onClick={() => updateLocation()} ><Icon.Crosshair /> Standort bestimmen</Button>
                        { showLocationSipnner &&
                            <div>
                                <Spinner animation="border" role="status">
                                    <span className="visually-hidden">Loading...</span>
                                </Spinner>
                            </div>
                        }
                    </div>
                    <div className='km-section'>
                        <Button className="km-btn-in-list" variant="primary" onClick={() => save()}><Icon.Save /> Speichern</Button>
                        <Button className="km-btn-in-list" variant="secondary" href="/list" ><Icon.SignStop /> Abbrechen</Button>
                    </div>
                </Form>
            </Container>
        </div>
    );
}
export default NewBox;