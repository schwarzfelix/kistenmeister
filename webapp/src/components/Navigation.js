import React from 'react';
import '../kistenmeister.css';
import * as Icon from 'react-bootstrap-icons';

function Navigation({ boxId, activeItem }) {


    function navigate(destination) {
        //setActiveItem(destination);
        window.location.href = "/box/" + boxId + "/" + destination;
    }

    function isActive(item) {
        return activeItem === item;
    }

    return (
        <div className="km-nav">
                <ul>
                    <li className={isActive("")         ? "km-nav-item-active" : "km-nav-item"} onClick={() => navigate("")}>
                        <div className='km-nav-icon'><Icon.CardHeading /></div>
                        <div className='km-nav-text'>Details</div>
                    </li>
                    <li className={isActive("comments") ? "km-nav-item-active" : "km-nav-item"} onClick={() => navigate("comments")}>
                        <div className='km-nav-icon'><Icon.Chat /></div>
                        <div className='km-nav-text'>Kommentare</div>
                    </li>
                    <li className={isActive("images") ? "km-nav-item-active" : "km-nav-item"} onClick={() => navigate("images")}>
                        <div className='km-nav-icon'><Icon.Image /></div>
                        <div className='km-nav-text'>Bilder</div>
                    </li>
                    <li className={isActive("qr")       ? "km-nav-item-active" : "km-nav-item"} onClick={() => navigate("qr")}>
                        <div className='km-nav-icon'><Icon.QrCode /></div>
                        <div className='km-nav-text'>QR-Code</div>
                    </li>
                </ul>
        </div>
    );
}
export default Navigation;