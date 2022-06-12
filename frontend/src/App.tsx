import type { Component } from 'solid-js';

import logo from './robot.png';
import styles from './App.module.css';

const App: Component = () => {
  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt="logo" />
        <p>
          Welcome to MadGo! Feel free to tinker with MadGo and Solidjs.
        </p>
        <a
          class={styles.link}
          href="https://github.com/solidjs/solid"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn Solid
        </a>
        <a
          class={styles.link}
          href="https://github.com/cybe42/MadGo"
          target="_blank"
          rel="noopener noreferrer"
        >
          MadGo Github
        </a>
      </header>
    </div>
  );
};

export default App;
