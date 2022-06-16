import React, { useState } from 'react'
import { useForm } from 'react-hook-form'
import { Container, Row, Col, Form, Button } from 'react-bootstrap'
import '../styles/Register.css'

export default function Register() {
    const [formRegist, setFormRegist] = useState(0)
    const { watch, register, handleSubmit, formState: { isValid } } = useForm({ mode: "all" })
    const completeFormRegist = () => {
        setFormRegist(cur => cur + 1)
    }
    const renderButton = () => {
        if (formRegist > 2) {
            return undefined
        } else if (formRegist === 2) {
            return (
                <Button
                    variant="primary" type="submit" disabled={!isValid}>
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
        window.alert(JSON.stringify(values, null, 2))
        completeFormRegist();
    }
    return (
        <Container className="container-regist">
            <Row>
                <Col className="apa">
                    <figure className="position-relative">
                        <img src={require("../images/book.png")} className="img-fluid"></img>
                        <figcaption>
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
                                                <p>Sudah punya akun PinjamBuku? <a href="/login">Masuk</a></p> <br />
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
                                                <h1>Isi Data Diri</h1><br />
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

                                        {formRegist >= 2 && (
                                            <section style={{ display: formRegist === 2 ? "block" : "none" }}>
                                                <h1>Isi Data Diri</h1><br />
                                                <Form.Group className="mb-3" controlId="formBasicNoKTP">
                                                    <Form.Label>No. KTP</Form.Label>
                                                    <Form.Control type="number" name="no_ktp" {...register("no_ktp", { required: { value: true } })} />
                                                </Form.Group>

                                                <Form.Group className="mb-3" controlId="formBasicPhotoKTP">
                                                    <Form.Label>Upload Photo KTP</Form.Label>
                                                    <Form.Control type="file" name="picture_ktp" {...register("picture_ktp", { required: { value: true } })} />
                                                </Form.Group>
                                            </section>
                                        )}

                                        {formRegist === 3 && (
                                            <section>
                                                <h1 className="alert-title">Berhasil Terdaftar!</h1>
                                                <p className="alert-words">Silahkan menunggu konfirmasi selanjutnya terkait verifikasi akun melalui email Anda. <br /><span>Terima kasih.</span></p>
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
    )
}