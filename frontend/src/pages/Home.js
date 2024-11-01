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
                        <h2>{profile.name}</h2>
                        <p>{profile.bio}</p>
                        <img src={profile.profile_picture} alt={profile.name} />

                        {/* Tampilkan Projects jika ada */}
                        <h3>Projects:</h3>
                        <ul>
                            {profile.projects && profile.projects.map((project, idx) => (
                                <li key={idx}>{project.name} - {project.description}</li>
                            ))}
                        </ul>

                        {/* Tampilkan Skills jika ada */}
                        <h3>Skills:</h3>
                        <ul>
                            {profile.skills && profile.skills.map((skill, idx) => (
                                <li key={idx}>{skill.name}</li>
                            ))}
                        </ul>

                        {/* Tambahkan bagian lain sesuai struktur JSON */}
                    </li>
                ))}
            </ul>
        </div>
    );

};

export default Portfolio;
