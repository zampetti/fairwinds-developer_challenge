import React from 'react';
import Board from './components/board'
import styles from './App.scss';

type AppProps = { title: string };
 
const App = ({ title }: AppProps) => (
  <div>
    <div className={styles.body}>
      <h3>{title}</h3>
      <Board />
      {/* <div className={styles.issues}>
        <Issue issue={"immigration"} />
        <Issue issue={"#BlackLivesMatter"} />
        <Issue issue={"climage change"} />
        <Issue issue={"immigration"} />
        <Issue issue={"#BlackLivesMatter"} />
        <Issue issue={"climage change"} />
      </div> */}
    </div>
    <div className={styles.footer}></div>
  </div>
)
 
export default App;