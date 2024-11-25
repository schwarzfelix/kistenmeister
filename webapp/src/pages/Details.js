import React from 'react';
import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import * as Icon from 'react-bootstrap-icons';
import { Table, Form, Button, Spinner, Tab, Alert, Container } from 'react-bootstrap';
import Navigation from '../components/Navigation';
import Header from '../components/Header';
import Config from '../km-config';
import { useContext } from 'react';
import { UserContext } from '../App';

function Details() {

    const { id } = useParams()
    const [result , setResult] = useState('');
    const [editMode, setEditMode] = useState(false);
    const [showLocationSipnner, setShowLocationSpinner] = useState(false);

    const [showAlert, setShowAlert] = useState(false);
    const [alertVariant, setAlertVariant] = useState('');
    const [alertText, setAlertText] = useState('');

    const [resultFavs , setResultFavs] = useState();
    const [isFavState, setIsFavState] = useState(false);

    const { user, setUser } = useContext(UserContext);

    const formatDateTime = (datetime) => {
        const date = new Date(datetime);
        const formattedDate = date.toLocaleDateString('de-DE'); // Formats to 'MM.DD.YYYY'
        const formattedTime = date.toLocaleTimeString('de-DE'); // Formats to 'HH:MM:SS 
        return `${formattedDate} ${formattedTime}`;
    };

    let updateIsFavState = () => {
        console.log("Update Fav State");
        try {
            let favs = resultFavs["Merklisteneinträge"];
            let favsIds = favs.map(fav => fav.Kiste_id);
            if (favsIds.includes(Number(id))) {
                setIsFavState(true);
            } else {
                setIsFavState(false);
            }
        } catch (error) {
            setIsFavState(false);
        }
        console.log(isFavState);
    }

    function getData() {
        try {
            fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kiste/' + id,
            {
                method: "GET",
                headers: { 'Authorization': ("Bearer " + user.token) }
            }
            )
                .then(response => response.json())
                .then(result => setResult(result));
        } catch (error) {
            console.log(error);
            setTimeout(getData, 5000);
        }
    }

    function toggleEditMode() {
        setEditMode(!editMode);
        setShowAlert(false);
        setShowLocationSpinner(false);
    }

    function saveData() {
        console.log("Save Data");
        setShowAlert(false)
        setShowLocationSpinner(false);


        const data = new FormData();
        data.append("Name", document.getElementById('name').value ? document.getElementById('name').value : result.Kiste.Name);
        data.append("Beschreibung", document.getElementById('beschreibung').value ? document.getElementById('beschreibung').value : result.Kiste.Beschreibung);
        data.append("Verantwortlicher", document.getElementById('verantwortlicher').value ? document.getElementById('verantwortlicher').value : result.Kiste.Verantwortlicher);
        data.append("Ort", document.getElementById('ort').value ? document.getElementById('ort').value : result.Kiste.Ort);

        console.log(data);

        try {
            fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kiste/' + id,
            {
                method: "PUT",
                headers: { 'Authorization': ("Bearer " + user.token) },
                body: data
            }
            )
                .then(response => response.json())
                .then(result => setResult(result));
        } catch (error) {
            console.log(error);
        }

        window.location.reload();
    }

    function deleteBox() {
        console.log("Delete Box");
        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kiste/' + id,
        {
            method: "DELETE",
            headers: { 'Authorization': ("Bearer " + user.token) }
        }
        )
            .then(response => response.text())
            .then(result => {
                console.log(result);
                window.location.href = '/list';
            });
    }

    let fetchFavs = () => {

        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/merklisteneinträge',
            {
                method: "GET",
                headers: { 'Authorization': ("Bearer " + user.token) }
            }
            )
            .then(response => response.json())
            .then(result => {
                setResultFavs(result);
            }
        );
    }

    let addFav = (id) => {
        const data = new FormData();
        data.append("kiste_id", id);

        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/merklisteneintrag',
            {
                method: "POST",
                headers: { 'Authorization': ("Bearer " + user.token) },
                body: data
            }
            )
            .then(response => response.text())
            .then(result => {
                window.location.reload();
            }
        );
    }

    let removeFav = () => {
        let merklisteneintrag_id = resultFavs["Merklisteneinträge"].find(fav => fav.Kiste_id === Number(id)).ID;

        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/merklisteneintrag/' + merklisteneintrag_id,
            {
                method: "DELETE",
                headers: { 'Authorization': ("Bearer " + user.token) }
            }
            )
            .then(response => response.text())
            .then(result => {
                window.location.reload();
            }
        );
    }


    function updateLocation() {
        setShowLocationSpinner(true);
        navigator.geolocation.getCurrentPosition((position) => {

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

    useEffect(() => {
        getData();
        fetchFavs();
    }, []);

    useEffect(() => {
        updateIsFavState();
    }, [resultFavs]);

    return (
        <div>
            <Header />
            <Navigation activeItem="" boxId={id} />
            <Container className='km-page-content'>
            { !result &&
                <div>
                    <Spinner animation="border" role="status">
                        <span className="visually-hidden">Loading...</span>
                    </Spinner>
                </div>
            }
            {result && result.error &&
                <div>
                    <Alert variant="error">
                        {result.error}
                    </Alert>
                </div>
            }
            {result && result.Kiste &&
                <div>
                    
                    <Form>

                        { showAlert &&

                            <div className='km-section'>
                                <Alert variant={alertVariant}>
                                    {alertText}
                                </Alert>
                            </div>
                        }

                        <div className='km-section'>
                            {editMode &&
                                <div>
                                    <div className='km-form-space'>
                                        <Form.Group controlId="name">
                                            <Form.Control type="text" size="lg" placeholder={result.Kiste.Name} />
                                        </Form.Group>
                                    </div>
                                    <div className='km-form-space'>
                                        <Form.Group controlId="beschreibung">
                                            <Form.Control as="textarea" rows={3} placeholder={result.Kiste.Beschreibung} />
                                        </Form.Group>
                                    </div>
                                </div>
                            }
                            {!editMode && 
                                <div>
                                    <h1>{result.Kiste.Name}</h1>
                                    <div>{result.Kiste.Beschreibung}</div>
                                </div>
                            }
                        </div>

                        { !editMode && resultFavs &&
                            <div className='km-section'>
                                {
                                isFavState
                                ?
                                <Button className="km-btn-in-list" variant="secondary" onClick={() => removeFav(id)}><Icon.StarFill /> Von Liste entfernen</Button>
                                :
                                <Button className="km-btn-in-list" variant="primary" onClick={() => addFav(id)}><Icon.Star /> Auf die Merkliste</Button>
                                }
                                <Button className="km-btn-in-list" variant="secondary" onClick={() => setEditMode(!editMode)}><Icon.Pencil /> Bearbeiten</Button>
                            </div>
                        }

                        <div className='km-section'>
                            <Table>
                                <tbody>
                                    <tr>
                                        <td><Icon.Person /> Verantwortlich:</td>
                                        <td>
                                            {editMode && 
                                                <Form.Group controlId="verantwortlicher">
                                                    <Form.Control type="text" size="sm" placeholder={result.Kiste.Verantwortlicher} />
                                                </Form.Group>
                                            }
                                            {!editMode && 
                                                result.Kiste.Verantwortlicher
                                            }
                                        </td>
                                    </tr>
                                    <tr>
                                        <td><Icon.PinMap /> Ort: </td>
                                        <td>
                                            {editMode &&
                                            <div>
                                                <Form.Group controlId="ort">
                                                    <Form.Control type="text" size="sm" placeholder={result.Kiste.Ort} />
                                                </Form.Group>
                                                <Button size='sm' variant="light" onClick={() => updateLocation()}><Icon.Crosshair /> Standort bestimmen</Button>
                                                { showLocationSipnner &&
                                                    <div>
                                                        <Spinner animation="border" role="status">
                                                            <span className="visually-hidden">Loading...</span>
                                                        </Spinner>
                                                    </div>
                                                }
                                            </div>
                                            }
                                            {!editMode && 
                                                <a href={"https://www.google.com/maps/search/" + result.Kiste.Ort} target="_blank">{result.Kiste.Ort}</a>
                                            }
                                        </td>
                                    </tr>
                                </tbody>
                            </Table>
                        </div>

                        { editMode &&
                            <div className='km-section'>
                                <Button className="km-btn-in-list" variant="primary" onClick={() => saveData()}><Icon.Save /> Speichern</Button>
                                <Button className="km-btn-in-list" variant="danger" onClick={() => deleteBox()}><Icon.Trash /> Kiste Löschen</Button>
                                <Button className="km-btn-in-list" variant="secondary" onClick={() => toggleEditMode()}><Icon.SignStop /> Abbrechen</Button>
                            </div>
                        }
                        
                        { !editMode &&
                            <div className='km-section'>
                                <hr />
                                <Table>
                                    <thead>
                                        <tr>
                                            <th>Erstellt</th>
                                            <th>Zuletzt geändert</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td>{formatDateTime(result.Kiste.Erstellungsdatum)}</td>
                                            <td>{formatDateTime(result.Kiste.Änderungsdatum)}</td>
                                        </tr>
                                        <tr>
                                            <td>{result.Kiste.Ersteller}</td>
                                            <td>{result.Kiste.Änderer}</td>
                                        </tr>
                                    </tbody>
                                </Table>
                            </div>
                        }
                        </Form>
                </div>
            }
            </Container>
        </div>
    );
}
export default Details;