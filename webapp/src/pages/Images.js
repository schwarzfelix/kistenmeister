import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Spinner, Button, Form, Container } from 'react-bootstrap';
import Navigation from '../components/Navigation';
import Header from '../components/Header';
import '../kistenmeister.css';
import * as Icon from 'react-bootstrap-icons';
import Config from '../km-config';
import Image from '../components/Image';
import { useContext } from 'react';
import { UserContext } from '../App';

function Images() {

    const { id } = useParams()
    const [result , setResult] = useState('');

    const [selectedImage, setSelectedImage] = useState(null);

    const { user, setUser } = useContext(UserContext);

    function uploadImage() {
        const file = document.getElementById('file').files[0];
        const data = new FormData();
        data.append("bild", file);
        data.append("Ersteller", "Felix");

        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/bild/' + id,
            {
                method: "POST",
                headers: { 'Authorization': ("Bearer " + user.token) },
                body: data
            }
        )
            .then(response => response.text())
            .then(text => {
                toObjectResult(text);
                window.location.reload();
            });
    }

    useEffect(() => {
        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/bilder/' + id,
        {
            method: "GET",
            headers: { 'Authorization': ("Bearer " + user.token) }
        }
        )
            .then(response => response.text())
            .then(text => toObjectResult(text));
    }, []);

    function toObjectResult(text) {
        let result = "{ \"Bilder\": " + text + "}";
        let json = JSON.parse(result);
        console.log(json);
        setResult(json);
    }

    return (
        <div>
            <Header />
            <Navigation activeItem="images" boxId={id} />
            <Container className='km-page-content'>
            <Form>
                <h6>Neues Bild hochladen:</h6>
                <div className="km-form-space">
                    <input
                        type="file"
                        id="file"
                    />
                </div>
                <div className="km-form-space">
                    <Button variant='primary' onClick={() => uploadImage()} ><Icon.Upload /> Bild hochladen</Button>
                </div>
            </Form>
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
            {result && result.Bilder && result.Bilder.length > 0 &&
                <div>
                    {result.Bilder.slice(0).reverse().map((image, index) => (
                        <Image id={image.ID} imagedata={image.Bild} author={image.Ersteller} date={image.Erstellungsdatum} />
                    ))}
                </div>
            }
            </Container>
        </div>
    );
}
export default Images;