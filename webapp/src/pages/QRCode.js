import React from 'react';
import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import {QRCodeSVG} from 'qrcode.react';
import { Card, CardBody, CardTitle, CardImg, Button, Spinner, Alert, Container } from 'react-bootstrap';
import Navigation from '../components/Navigation';
import Header from '../components/Header';
import * as Icon from 'react-bootstrap-icons';
import Config from '../km-config';
import { useContext } from 'react';
import { UserContext } from '../App';

function QRCode() {

    const { id } = useParams()
    const [result , setResult] = useState('');
    const [showAlert, setShowAlert] = useState(false);

    const { user, setUser } = useContext(UserContext);

    useEffect(() => {
        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kiste/' + id,
        {
            method: "GET",
            headers: { 'Authorization': ("Bearer " + user.token) }
        }
        )
            .then(response => response.json())
            .then(result => setResult(result));
    }, []);

    return (
        <div>
            <Header />
            <Navigation activeItem="qr" boxId={id} />
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
                        <p>Error: {result.error}</p>
                    </div>
                }
                {result && result.Kiste &&
                    <div className='km-qrcard'>
                        { showAlert &&
                            <Alert variant="success">
                                Link wurde in die Zwischenablage kopiert!
                            </Alert>
                        }
                        <Card>
                            <CardBody>
                                <QRCodeSVG value={"http:localhost:3000/box/" + id} />
                                <div className='km-form-space'>
                                    <CardTitle>{result.Kiste.Name}</CardTitle>
                                </div>
                                <Button variant="primary" onClick={() => {
                                    navigator.clipboard.writeText(Config.webapp.protocol + "://" + Config.webapp.host + ":" + Config.webapp.port + "/box/" + id);
                                    setShowAlert(true);
                                }}><Icon.Copy /> Link kopieren</Button>
                            </CardBody>
                        </Card>
                    </div>
                }
            </Container>
        </div>
    );
}
export default QRCode;