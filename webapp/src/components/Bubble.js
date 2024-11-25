import React from 'react';
import { Button } from 'react-bootstrap';
import '../kistenmeister.css';
import * as Icon from 'react-bootstrap-icons';
import Config from '../km-config';
import { useContext } from 'react';
import { UserContext } from '../App';

function Bubble({ id, content, author, date }) {

    const { user, setUser } = useContext(UserContext);

    const formatDateTime = (datetime) => {
        const date = new Date(datetime);
        const formattedDate = date.toLocaleDateString('de-DE'); // Formats to 'MM.DD.YYYY'
        const formattedTime = date.toLocaleTimeString('de-DE'); // Formats to 'HH:MM:SS 
        return `${formattedDate} ${formattedTime}`;
    };

    function deleteComment() {
        console.log("Delete Comment");
        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kommentar/' + id,
        {
            method: "DELETE",
            headers: { 'Authorization': ("Bearer " + user.token)}
        }
        )
            .then(response => response.json())
            .then(result => {
                console.log(result);
                window.location.reload();
            });
    }

    return (
        <div className="km-bubble">
            <div className="km-bubble-header">
                
            </div>
            <div className="km-bubble-inside">
                <div className="km-bubble-content">
                    {content}
                </div>
                <div className="km-bubble-meta">
                    { id &&
                        <Button className="km-btn-in-list" size="sm" onClick={() => deleteComment()}><Icon.Trash /></Button>
                    }
                {author} {date ? formatDateTime(date) : ""}
                </div>
            </div>
        </div>
    );
}
export default Bubble;