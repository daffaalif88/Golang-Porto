// frontend/src/Portfolio.js
import React, { useEffect, useState } from 'react';

const Portfolio = () => {
    const [data, setData] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/portofolio')
            .then((response) => response.json())
            .then((data) => {
                console.log(data); // Cek struktur data di konsol
                setData(data);
            })
            .catch((error) => console.error("Error fetching data:", error));
    }, []);


    return (
        <div>
            <h1>Portofolio</h1>
            <ul>
                {data.map((profile, index) => (
                    <li key={index}>
                        <h2>{profile.Name}</h2>
                        <p>{profile.Bio}</p>
                        <p>{profile.BirthDate}</p>
                        <img src={profile.profile_picture} alt={profile.name} />
                    </li>
                ))}
            </ul>
        </div>
    );

};

export default Portfolio;
