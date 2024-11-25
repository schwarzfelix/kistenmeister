import React from 'react';
import * as Icon from 'react-bootstrap-icons';
import '../kistenmeister.css'; // Ensure this CSS file includes the necessary styles

function Aktivierung() {
    return (
        <div className="App-aktivierung">
            <h1 className="text-center text-aktivierung">
                <Icon.PersonPlus className="me-2" />
                Genehmigen Sie den Link, den zur genannten E-Mail Adresse geschickt wurde!
            </h1>
        </div>
    );
}

export default Aktivierung;
