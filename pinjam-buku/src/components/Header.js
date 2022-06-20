import React from 'react'
import { Link } from "react-router-dom"
import { Navbar, Container, Nav, Button } from 'react-bootstrap'
import '../styles/Header.css'

export default function Header() {
    return (
        <Navbar>
            <Container className="nav">
                <Navbar.Brand href="/" className='logo' component={Link}>
                    PINJAMBUKU
                </Navbar.Brand>
                <Navbar.Toggle />
                <Navbar.Collapse className="justify-content-end">
                    <Nav.Link href="/galeri-buku" component={Link}>Galeri Buku</Nav.Link>
                    <Nav.Link href="/bantuan" component={Link}>Bantuan</Nav.Link>
                    <Nav.Link href="/kontak" component={Link}>Kontak</Nav.Link>
                    <Nav.Link href="/masuk" component={Link}><Button className="login">Masuk</Button></Nav.Link>
                    <Nav.Link href="/daftar" component={Link}><Button className="signup">Daftar</Button></Nav.Link>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    )
}