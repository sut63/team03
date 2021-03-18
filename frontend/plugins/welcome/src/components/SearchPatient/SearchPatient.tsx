import React, { useState } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import moment from 'moment';
import { Page, pageTheme, Header, Content} from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { EntDentalexpense, EntPatient } from '../../api';
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

  const [patient, setPatient] = useState<EntPatient[]>([])
  const [patientID, setPatientID] = useState(String);
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);




  const PatientIDhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setPatientID(event.target.value as string);
  };

  const cleardata = () => {
    setPatientID("");
    setStatus(false);
    setPatient([]);

  }

  

  const SearchPatient = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchpatients?patient=${patientID}`;
    const requestOptions = {
        method: 'GET',
    };
    fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
            console.log(data.data)
            setStatus(true);
            setAlert(false);
            setPatient([]);

          if (data.data != null) {
                if(data.data.length >= 1) {
                  setStatus(true);
                  setAlert(true);
                  console.log(data.data)
                  setPatient(data.data);
                }
            }
        });

}

  return (

    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ค้นหาประวัติผู้ป่วย">
        <table>
        <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
        </Button>
        </table>

      </Header>
      <Content>
        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>

              <Typography align="center" >
                <div style={{ background: "#1976d2", height: 60 }}>
                  <h1 style={
                    {
                      color: "#FFFFFF",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '30px',
                    }}>
                    ค้นหาประวัติผู้ป่วย
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>รหัสประจำตัวผู้ป่วย EX.P000001</div>
                    <TextField
                      id="patientID"
                      value={patientID}
                      onChange={PatientIDhandlehange || ''}
                      type="string"
                      size="medium"

                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>

                <div></div>
                
                <Button
                  onClick={() => {
                    SearchPatient();
                  }}

                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#00b0ff", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 10px',

                      }
                    }>
                    Search
            </h3>
                </Button>
                <Button
                  onClick={() => {
                    cleardata();

                  }}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#f9a280", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#000000",
                        padding: '0 25px',

                      }
                    }>
                    Clear
            </h3>
                </Button>
              </Typography>
            </Paper>

            {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    แสดงประวัติผู้ป้วย
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
              <TableCell align="center">PatientID</TableCell>
              <TableCell align="center">Name</TableCell>
              <TableCell align="center">Age</TableCell>
              <TableCell align="center">Disease</TableCell>
              <TableCell align="center">MedicalCare</TableCell>
              <TableCell align="center">Tel.</TableCell>
              <TableCell align="center">BirthDay</TableCell>
              </TableRow>
            </TableHead>
          <TableBody>
            {patient.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.PatientID}</TableCell>
                    <TableCell align="center">{item.Name}</TableCell>
                    <TableCell align="center">{item.Age}</TableCell>
                    <TableCell align="center">{item.edges?.Disease?.name}</TableCell>
                    <TableCell align="center">{item.edges?.Medicalcare?.name}</TableCell>
                    <TableCell align="center">{item.Tel}</TableCell>
                    <TableCell align="center">{moment(item.Birthday).format('DD/MM/YYYY')}</TableCell>
                                    
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