import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Spinner, Container, Form, Button } from 'react-bootstrap';
import Navigation from '../components/Navigation';
import Header from '../components/Header';
import Bubble from '../components/Bubble';
import '../kistenmeister.css';
import * as Icon from 'react-bootstrap-icons';
import Config from '../km-config';
import { useContext } from 'react';
import { UserContext } from '../App';

function Comments() {

    const { id } = useParams()
    const [result , setResult] = useState('');

    const { user, setUser } = useContext(UserContext);

    const formatDateTime = (datetime) => {
        const date = new Date(datetime);
        const formattedDate = date.toLocaleDateString('de-DE'); // Formats to 'MM.DD.YYYY'
        const formattedTime = date.toLocaleTimeString('de-DE'); // Formats to 'HH:MM:SS 
        return `${formattedDate} ${formattedTime}`;
    };

    useEffect(() => {
        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kommentare/' + id,
        {
            method: "GET",
            headers: { 'Authorization': ("Bearer " + user.token) }
        }
        )
            .then(response => response.json())
            .then(result => setResult(result));
    }, []);

    function saveComment() {

        console.log("Save Comment");
        const data = new FormData();
        data.append("Kommentar", document.getElementById('message').value);
        data.append("Ersteller", "Anonymous");
        console.log(data);

        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kommentar/' + id,

        {
            method: "POST",
            headers: { 'Authorization': ("Bearer " + user.token) },
            body: data
        }
        )
            .then(response => response.json())
            //.then(result => setResult(result));
            .then(result => {
                console.log(result);
                window.location.reload();
            });
    }


    return (
        <div>
            <Header />
            <Navigation activeItem="comments" boxId={id} />
            <Container className='km-page-content'>
            <Bubble content={
                <Form>
                    <div className='km-form-space'>
                        <Form.Group controlId="message">
                            <Form.Control as="textarea" rows={1} placeholder="Neuer Kommentar" />
                        </Form.Group>
                    </div>
                    <div className='km-form-space'>
                        <Button variant="light" onClick={() => saveComment()}><Icon.Send /> Speichern</Button>
                    </div>
                </Form>
            } />
            <hr />
            { !result &&
                <div>
                    <Spinner animation="border" role="status">
                        <span className="visually-hidden">Loading...</span>
                    </Spinner>
                </div>
            }
            {result && result.Error &&
                <div>
                    <p>Error: {result.Error}</p>
                </div>
            }
            {result && result.Kommentare &&
                <div>
                        {result.Kommentare.slice(0).reverse().map((kommentar, index) => (
                                <Bubble  id={kommentar.ID} content={kommentar.Kommentar} author={kommentar.Ersteller} date={kommentar.Erstellungsdatum} />
                        ))} 
                </div>
            }
            </Container>
        </div>
    );
}
export default Comments;