import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { EntMedicalfile } from '../../api';
import { Alert } from '@material-ui/lab';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    headsearch: {
      width: 'auto',
      margin: '10px',
      color: '#FFFFFF',
      background: '#2196F3',
    },
    margin: {
      margin: theme.spacing(1),
    },
    margins: {
      margin: theme.spacing(2),
    },

    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
    },
    table: {
      minWidth: 500,
    },


  }),
);


export default function ComponentsTable() {
  const classes = useStyles();
  const [medno, setMedno] = useState(String);
  const [medicalfile, setMedicalfile] = useState<EntMedicalfile[]>([]);
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  
  const Mednohandlechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setMedno(event.target.value as string);
  };

  const cleardata = () => {
    setMedno("");
    setStatus(false);
    setMedicalfile([]);

  }

  const SearchMedicalfile = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchmedicalfiles?medicalfile=${medno}`;
    const requestOptions = {
        method: 'GET',
    };
    fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
            console.log(data.data)
            setStatus(true);
            setAlert(false);
            setMedicalfile([]);

          if (data.data != null) {
                if(data.data.length >= 1) {
                  setStatus(true);
                  setAlert(true);
                  console.log(data.data)
                  setMedicalfile(data.data);
                }
            }
        });

}




  return (

    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ค้นหาประวัติทันตกรรมของผู้ป่วย">
        <table>
          <tr>
           
            <th>
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          <Link component={RouterLink} to="/Menu">
                <Button variant="contained" style={{ height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#000000",
                        borderRadius: 10,
                        height: 25,
                        padding: '0 20px',
                      }
                    }>
                    Back
            </h3>
                </Button>
              </Link>
            </th>
          </tr>
        </table>

      </Header>
      <Content>
        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>

              <Typography align="center" >
                <div style={{ background: "#3366CC", height: 50 }}>
                  <h1 style={
                    {
                      color: "#FFFFFF",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '30px',
                    }}>
                    ค้นหาประวัติทันตกรรม
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>หมายเลขทันตกรรม ex.M001</div>
                    <TextField
                      id="medno"
                      value={medno}
                      onChange={Mednohandlechange}
                      type="string"
                      size="small"

                      style={{ width: 200 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    SearchMedicalfile();

                  }}
                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#3399CC", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 10px',

                      }
                    }>
                    ค้นหาข้อมูล
            </h3>
                </Button>
                <Button
                  onClick={() => {
                    cleardata();

                  }}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#3399CC", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 25px',

                      }
                    }>
                    เคลียร์ข้อมูล
            </h3>
                </Button>
              </Typography>
            </Paper>

            {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    แสดงข้อมูลประวัติทันตกรรม
                  </Alert>
                )
                  : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่พบข้อมูล
                  </Alert>
                  )}
              </div>
            ) : null}

          </Grid>
        </Grid>


        <Grid  className={classes.paper}>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                          <TableCell align="center">Medicalfile_ID</TableCell>
                          <TableCell align="center">Medicalfile_No</TableCell>
                          <TableCell align="center">Patient</TableCell>
                          <TableCell align="center">Detail</TableCell>
                          <TableCell align="center">Dentist</TableCell>
                          <TableCell align="center">Date</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {medicalfile.map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.Medno}</TableCell>
                              <TableCell align="center">{item.edges?.Patient?.Name}</TableCell>
                              <TableCell align="center">{item.Detial}</TableCell>
                              <TableCell align="center">{item.edges?.Dentist?.name}</TableCell>
                              <TableCell align="center">{moment(item.AddedTime).format('DD/MM/YYYY HH:mm')}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  
         
        </Grid>
      </Content>
    </Page>
  );

}