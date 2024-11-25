import React from 'react';
import { useContext } from 'react';
import { Navbar, Container, Nav, NavDropdown } from 'react-bootstrap';
import '../kistenmeister.css';
import * as Icon from 'react-bootstrap-icons';
import { UserContext } from '../App';

function Header({ boxId, activeItem }) {

    const { user, setUser } = useContext(UserContext);

    return (
        <div className="km-header">
            <Navbar expand="lg" className="bg-body-tertiary">
                <Container>
                    <Navbar.Brand href="/list?onlyFavs=false">
                        <Icon.BoxSeam className="me-2" />
                        Kistenmeister
                    </Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav" className="justify-content-end">
                        <Nav.Link href="/list?onlyFavs=true">
                            <div className="km-header-link"><Icon.ListStars /> Merkliste</div>
                        </Nav.Link>
                        <Nav.Link href="/new">
                            <div className="km-header-link"><Icon.Plus /> Neue Kiste</div>
                        </Nav.Link>
                        <Nav.Link href="/profile">
                            <div className="km-header-link"><Icon.Person /> {user.name}</div>
                        </Nav.Link>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        </div>
    );
}
export default Header;