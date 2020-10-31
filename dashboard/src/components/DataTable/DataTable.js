import React from 'react';
import styles from './DataTable.module.css';


function DataTable({ tableData}) {

    const formatTableData = () => {
        let tableCells;
        const tableRow = []; 

        console.log("tabledata",tableData)
        tableData.forEach( data => {
            console.log("data",data)
            tableCells = []
            Object.values(data).forEach ( rowData => {
                console.log("rowdata",rowData)
                for( let item in rowData ){
                    tableCells.push(
                        <tr className={styles['tr']}>
                            <td className={styles['td']}>{item}</td>
                            <td className={styles['td']}>{rowData[item]}</td>
                        </tr>
                    )
                }
            }
        )
        tableRow.push(
            <table>
                <thead className={styles['thead']}>
                    <tr></tr>{Object.keys(data)}</thead>
                <tbody>{tableCells}</tbody>
            </table>
       )
    })

    return tableRow;

    }

    return(
        <div className={styles['table']}>
            {formatTableData()}
        </div>
    );
}

export default DataTable;