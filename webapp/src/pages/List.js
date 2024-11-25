import React, { useState, useEffect, useContext } from 'react';
import { Button, Table, Spinner, Container } from 'react-bootstrap';
import Dropdown from 'react-bootstrap/Dropdown';
import * as Icon from 'react-bootstrap-icons';
import Header from '../components/Header';
import Config from '../km-config';
import { useSearchParams } from 'react-router-dom';
import { UserContext } from '../App';

function List() {

    const [onlyFavs , setOnlyFavs] = useState(false);
    const [resultAll , setResultAll] = useState();
    const [resultFavs , setResultFavs] = useState();

    const [listData , setListData] = useState();

    const [searchParams, setSeachParams] = useSearchParams( );
    const { user, setUser } = useContext(UserContext);

    let updateListData = () => {

        try {
            if (onlyFavs) {
                let favs = resultFavs["Merklisteneinträge"];
                let all = resultAll["Alle Kisten"];
    
                let favsIds = favs.map(fav => fav.Kiste_id);
                let list = all.filter(box => favsIds.includes(box.ID));
                setListData(list);
            } else {
                setListData(resultAll["Alle Kisten"]);
            }
        } catch (error) {
            setListData([]);
            console.error("Error in updateListData: " + error);
        }
    }

    let isFav = (id) => {

        try {

            let favs = resultFavs["Merklisteneinträge"];
            let favsIds = favs.map(fav => fav.Kiste_id);
            return favsIds.includes(id);

        } catch (error) {
            console.error("Error in isFav: " + error);
            return false;
        }
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

    let removeFav = (id) => {
        let merklisteneintrag_id = resultFavs["Merklisteneinträge"].find(fav => fav.Kiste_id === id).ID;

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

    let fetchBoxes = () => {

        fetch(Config.server.protocol + "://" + Config.server.host + ":" + Config.server.port + '/kisten',
            {
                method: "GET",
                headers: { 'Authorization': ("Bearer " + user.token) }
            }
            )
            .then(response => response.json())
            .then(result => {
                setResultAll(result);
            }
        );
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
                console.log("Favs:");
                console.log(result);
                setResultFavs(result);
            }
        );
    }

    useEffect(() => {

        fetchBoxes();
        fetchFavs();

        setOnlyFavs(searchParams.get("onlyFavs") === "true");

    }, []);

    useEffect(() => {
        setSeachParams({onlyFavs: onlyFavs});
    }, [onlyFavs]);

    // use effect with dependencies for the two fetches resultAll and resultFavs
    useEffect(() => {

        if (resultAll && resultFavs) {
            updateListData();
        }
    }, [resultAll, resultFavs, onlyFavs]);

    return (
        <div>
            <Header />
            <Container className='km-page-content'>
            <div className='km-section'>
                <h1>{onlyFavs ? "Merkliste" : "Alle Kisten"}</h1>
            </div>
            <div className='km-section'>
                <Button className="km-btn-in-list" variant='primary' href='/new'><Icon.Plus /> Neue Kiste</Button>
            </div>
            { !listData && <Spinner animation="border" /> }
            { listData && listData.length < 1 && <p>Keine Kisten gefunden</p> }
            {listData && listData.length > 0 &&
                <div>
                    <Table striped bordered hover>
                        <thead>
                            <tr>
                                <th>Kiste</th>
                                <th>Aktionen</th>
                            </tr>
                        </thead>
                        <tbody>
                            {listData.map((box, index) => (
                                <tr key={index}>
                                    <td>
                                        <Button variant="link" href={"/box/" + box.ID}>{box.Name}</Button>
                                    </td>
                                    <td>
                                        <Dropdown>
                                            <Dropdown.Toggle size="sm" variant="light" id="dropdown-basic">
                                                •••
                                            </Dropdown.Toggle>

                                            <Dropdown.Menu>
                                                <Dropdown.ItemText href={"/box/" + box.ID + "/qr"}>
                                                    { 
                                                    isFav(box.ID)
                                                    ?
                                                    <Button variant="light" size="sm" onClick={() => removeFav(box.ID)}><Icon.StarFill /> Von Liste entfernen</Button>
                                                    :
                                                    <Button variant="light" size="sm" onClick={() => addFav(box.ID)}><Icon.Star /> Auf die Merkliste</Button>
                                                    }
                                                </Dropdown.ItemText>
                                                <hr />
                                                <Dropdown.Item href={"/box/" + box.ID}><Icon.CardHeading /> Details</Dropdown.Item>
                                                <Dropdown.Item href={"/box/" + box.ID + "/comments"}><Icon.Chat /> Kommentare</Dropdown.Item>
                                                <Dropdown.Item href={"/box/" + box.ID + "/images"}><Icon.Image /> Bilder</Dropdown.Item>
                                                <Dropdown.Item href={"/box/" + box.ID + "/qr"}><Icon.QrCode /> QR-Code</Dropdown.Item>
                                            </Dropdown.Menu>
                                        </Dropdown>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </Table>
                </div>
            }
            </Container>
        </div>
    );
}
export default List;