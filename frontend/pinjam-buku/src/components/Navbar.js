import React from 'react'
import { Navbar, Container, Nav, Button } from 'react-bootstrap'
import "./Navbar.css"

export default function Header() {
    return (
        <Navbar className="bg-white">
            <Container>
                <Navbar.Brand href="#home" className='logo'>PINJAMBUKU</Navbar.Brand>
                <Navbar.Toggle />
                <Navbar.Collapse className="justify-content-end">
                    <Nav.Link href="/">Galeri Buku</Nav.Link>
                    <Nav.Link href="/">Bantuan</Nav.Link>
                    <Nav.Link href="/">Kontak</Nav.Link>
                    <Nav.Link href="/"><Button className="login">Masuk</Button></Nav.Link>
                    <Nav.Link href="/"><Button className="signup">Daftar</Button></Nav.Link>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    )
}