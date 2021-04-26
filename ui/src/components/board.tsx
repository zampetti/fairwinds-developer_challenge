import React from 'react';
// import { useDispatch, useSelector } from "react-redux";
import styles from './board.scss';
// import { sagaActions } from '../redux/sagaActions'
// import { selectFetch } from './issueSlice'

// type IssueProps = { issue: string };
 
const Board = () => {
    // const data = useSelector(selectFetch);
    // const dispatch = useDispatch();

    // console.log("DATA IN COMPONENT: ", data);
    
    return (
        <div className={styles.body}>
            <svg height="500" width="500">
                <line x1="175" y1="50" x2="175" y2="450" className={styles.stroke} />
                <line x1="325" y1="50" x2="325" y2="450" className={styles.stroke} />
                <line x1="50" y1="175" x2="450" y2="175" className={styles.stroke} />
                <line x1="50" y1="325" x2="450" y2="325" className={styles.stroke} />
            </svg>
        </div>
    )
}
 
export default Board;