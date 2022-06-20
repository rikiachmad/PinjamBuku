import React, { useState } from 'react'
import Header from '../components/Header'
import { useNavigate } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { Container, Row, Col, Form, Button } from 'react-bootstrap'
import { FaArrowLeft } from 'react-icons/fa'
import '../styles/user/Register/Register.css'
import Swal from 'sweetalert2'

export default function Register() {
    const navigate = useNavigate();
    const [formRegist, setFormRegist] = useState(0)
    const { watch, register, handleSubmit, formState: { isValid } } = useForm({ mode: "all" })
    const completeFormRegist = () => {
        setFormRegist(cur => cur + 1)
    }
    const backFormRegist = () => {
        setFormRegist(cur => cur - 1)
    }
    const popupAlert = () => {
        Swal.fire("Berhasil!", "Akun Anda telah terdaftar.", "success");
        navigate("/login")
    }
    const renderButton = () => {
        if (formRegist > 1) {
            return undefined
        } else if (formRegist === 1) {
            return (
                <Button
                    variant="primary" type="submit" disabled={!isValid} onClick={popupAlert}>
                    Daftar
                </Button>)
        } else {
            return (
                <Button onClick={completeFormRegist}
                    variant="primary" type="button" disabled={!isValid}>
                    Lanjutkan
                </Button>)
        }
    }
    const submitForm = (values: any): void => {
        // testing
        // window.alert(JSON.stringify(values, null, 2))
        completeFormRegist();
    }
    return (
        <>
        <Header />
            <Container className="container-regist">
                <Row>
                    <Col className="logo-people">
                        <figure className="position-relative logo-regist">
                            <img src={require("../images/book.png")} className="img-fluid"></img>
                            <figcaption className="text-regist">
                                <span>PINJAMBUKU</span><br />
                                Platform peminjaman buku perpustakaan dari mana saja dengan mudah.
                            </figcaption>
                        </figure>
                    </Col>
                    <Col xs lg="6" className="registration">
                        <div className="jumbotron">
                            <Container>
                                <Row>
                                    <Col>
                                        <Form onSubmit={handleSubmit(submitForm)}>
                                            {formRegist >= 0 && (
                                                <section style={{ display: formRegist === 0 ? "block" : "none" }}>
                                                    <h1>Daftar Sekarang</h1><br />
                                                    <p>Sudah punya akun PinjamBuku? <a href="/login" class="login">Masuk</a></p> <br />
                                                    <Form.Group className="mb-3" controlId="formBasicEmail">
                                                        <Form.Label>Email</Form.Label>
                                                        <Form.Control type="email" name="email" {...register("email", { required: { value: true } })} />
                                                    </Form.Group>

                                                    <Form.Group className="mb-3" controlId="formBasicPassword">
                                                        <Form.Label>Password</Form.Label>
                                                        <Form.Control type="password" name="password" {...register("password", { required: { value: true } })} />
                                                    </Form.Group>
                                                </section>
                                            )}

                                            {formRegist >= 1 && (
                                                <section style={{ display: formRegist === 1 ? "block" : "none" }}>
                                                    <h1>
                                                        <Button className="btn-back" onClick={backFormRegist}><FaArrowLeft className="fa-left" /></Button> Isi Data Diri
                                                    </h1><br />
                                                    <Form.Group className="mb-3" controlId="formBasicName">
                                                        <Form.Label>Nama Lengkap</Form.Label>
                                                        <Form.Control type="text" name="fullname" {...register("fullname", { required: { value: true } })} />
                                                    </Form.Group>

                                                    <Form.Group className="mb-3" controlId="formBasicNoTlp">
                                                        <Form.Label>No. Telepon</Form.Label>
                                                        <Form.Control type="number" name="phone_number" {...register("phone_number", { required: { value: true } })} />
                                                    </Form.Group>

                                                    <Form.Group className="mb-3" controlId="formBasicAlamat">
                                                        <Form.Label>Alamat</Form.Label>
                                                        <Form.Control type="text" name="address" {...register("address", { required: { value: true } })} />
                                                    </Form.Group>
                                                </section>
                                            )}
                                            {renderButton()}
                                            {/* <pre>{JSON.stringify(watch(), null, 2)}</pre> */}
                                        </Form>
                                    </Col>
                                </Row>
                            </Container>
                        </div>
                    </Col>
                </Row>
            </Container>
        </>
    )
}