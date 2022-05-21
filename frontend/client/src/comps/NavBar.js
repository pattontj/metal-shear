import React, { Component } from 'react';

import { Link } from 'react-router-dom';
import style from './NavBar.css'

class NavBar extends React.Component {
    render() { 
        return (
            <nav>
                <Link to="/" className="link">Home</Link>
                <Link to="/statsuki" className="link">Statsuki </Link>
            </nav>
        );
    }
}

export default NavBar;