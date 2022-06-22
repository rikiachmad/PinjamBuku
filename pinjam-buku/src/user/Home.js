import React, { useState } from 'react'
import Header from '../components/Header';
import { useForm } from 'react-hook-form'
import { Container, Row, Col, Form, Button } from 'react-bootstrap'
import { MdOutlineTouchApp, MdOutlineLocalLibrary, MdPayment } from "react-icons/md";
import '../styles/user/Home/Home.css'

export default function Home() {
    return (
        <>
        <Header/>
            <Container className="landing-page">
                <section className="carousel">
                    <figure className="image-home">
                        <img src={require("../images/home.jpeg")} className="image-carousel img-fluid"></img>
                        <figcaption className="text-carousel">
                            PINJAM BUKU DENGAN MUDAH DIMANAPUN KAMU BERADA <br />
                        </figcaption>
                        <Button className="btn-search-home">Cari Buku Disini</Button>
                    </figure>
                </section>
                <section className="second row justify-content-center align-items-center w-100">
                    <h2 className="second-title">Membaca Buku Semakin Mudah</h2>
                    <p className="second-paragraph">
                        Baca buku, berbagi koleksi bacaan dan bersosialisasi secara bersamaan. 
                        Dimana pun, kapan pun dengan nyaman bersama setiap orang.
                    </p>
                </section>
                <section className="third d-flex justify-content-around">
                    <div className="third-benefit text-center">
                        <MdOutlineTouchApp className="rounded-circle" /><br/>
                        <p className="third-caption">Pinjam buku yang kamu inginkan dari  mana saja</p>
                    </div>
                    <div className="third-benefit text-center">
                        <MdOutlineLocalLibrary className="rounded-circle" /><br/>
                        <p className="third-caption">Bermacam - macam pilihan hingga 1000 buku</p>
                    </div>
                    <div className="third-benefit text-center">
                        <MdPayment className="rounded-circle" /><br/>
                        <p className="third-caption">Simpan uangmu, pinjam buku tanpa biaya pinjaman</p>
                    </div>
                </section>
                <section className="partner text-center">
                    <h1 className="partner-name">Partner Kami</h1>
                    <section className="partners-info d-flex justify-content-around">
                        <div className="partner-benefit text-center">
                            <img src={require("../images/library-logo.png")}></img>
                            <p className="partner-caption">Perpustakaan Provinsi Kalimantan Timur</p>
                        </div>
                        <div className="partner-benefit text-center">
                            <img src={require("../images/library-logo.png")}></img>
                            <p className="partner-caption">Perpustakaan Provinsi Kalimantan Timur</p>
                        </div>
                        <div className="partner-benefit text-center">
                            <img src={require("../images/library-logo.png")}></img>
                            <p className="partner-caption">Perpustakaan Provinsi Kalimantan Timur</p>
                        </div>
                        <div className="partner-benefit text-center">
                            <img src={require("../images/library-logo.png")}></img>
                            <p className="partner-caption">Perpustakaan Provinsi Kalimantan Timur</p>
                        </div>
                    </section>
                </section>
            </Container>
        </>
    )
}