import React from 'react';
import { Button } from 'react-bootstrap';
import '../kistenmeister.css';
import * as Icon from 'react-bootstrap-icons';
import Config from '../km-config';
import { useContext } from 'react';
import { UserContext } from '../App';

function Image({ id, imagedata, author, date }) {

    const { user, setUser } = useContext(UserContext);

    const formatDateTime = (datetime) => {
        const date = new Date(datetime);
        const formattedDate = date.toLocaleDateString('de-DE'); // Formats to 'MM.DD.YYYY'
        const formattedTime = date.toLocaleTimeString('de-DE'); // Formats to 'HH:MM:SS 
        return `${formattedDate} ${formattedTime}`;
    };

    function deleteImage(imageId) {
        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/bild/' + imageId,
        {
            method: "DELETE",
            headers: { 'Authorization': ("Bearer " + user.token)}
        }
        )
            .then(response => response.text())
            .then(result => {
                console.log(result);
                window.location.reload();
            });
    }

    return (
        <div className='km-image'>
            <img src={`data:image/jpeg;base64,${imagedata}`} />
            <div className='km-img-foot'>
                <Button className="km-btn-in-list" size="sm" variant='light' onClick={() => deleteImage(id)} ><Icon.Trash /></Button>
                {author} {date ? formatDateTime(date) : ""}
            </div>
        </div>
    );
}
export default Image;