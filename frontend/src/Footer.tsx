import React from 'react';

const Footer = () => {
  const currentYear = new Date().getFullYear();

  return (
    <footer style={styles.footer}>
      <p style={styles.text}>
        &copy; {currentYear} DecentralizedLottery. All rights reserved.
      </p>
      <ul style={styles.links}>
        <li><a href="/terms" style={styles.link}>Terms of Service</a></li>
        <li><a href="/privacy" style={styles.link}>Privacy Policy</a></li>
        <li><a href="/contact" style={styles.link}>Contact Us</a></li>
      </ul>
    </footer>
  );
};

const styles = {
  footer: {
    padding: '20px',
    textAlign: 'center',
    backgroundColor: '#282c34',
    color: '#fff',
    // marginTop: 'auto',
    marginTop: '40px',
  },
  text: {
    margin: 0,
    fontSize: '14px',
  },
  links: {
    listStyleType: 'none',
    padding: 0,
    margin: '10px 0 0',
    display: 'flex',
    justifyContent: 'center',
    gap: '15px',
  },
  link: {
    color: '#61dafb',
    textDecoration: 'none',
    fontSize: '14px',
  }
};

export default Footer;