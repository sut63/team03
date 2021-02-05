import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntPatient } from '../../api/models/EntPatient';
import Swal from 'sweetalert2'
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';


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
const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  //showCloseButton: true,

});


export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  //---------------------------
  const [checkpatientid, setcheckPatientids] = useState(false);
  const [patient, setPatient] = useState<EntPatient[]>([]);
  //--------------------------
  const [patientid, setPatientids] = useState(String);
  const profile = { givenName: 'ระบบค้นหาประวัติผู้ป่วย' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getPatient = async () => {
      const res = await api.listPatient({ offset: 0 });
      setLoading(false);
      setPatient(res);
    };
    getPatient();
  }, [loading]);

  //-------------------
  const handlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setcheckPatientids(false);
    setPatientids(event.target.value as string);

  };

  const cleardata = () => {
    setPatientids("");
    setSearch(false);
    setcheckPatientids(false);
    setSearch(false);

  }
  //---------------------
  const checkpatient = async () => {
    var check = false;
    patient.map(item => {
      if (patientid != "") {
        if (item.patientID?.includes(patientid)) {
          setcheckPatientids(true);
          alertMessage("success", "พบข้อมูล");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูล");
    }
    console.log(checkpatientid)
    if (patientid == "") {
      alertMessage("info", "แสดงข้อมูลประวัติผู้ป่วย");
    }
  };

  return (

    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ค้นหาประวัติผู้ป่วย">
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
                    <div className={classes.paper}><strong>หมายเลขประจำตัวผู้ป่วย EX.000001</strong></div>
                    <TextField
                      id="patientid"
                      value={patientid}
                      onChange={handlehange}
                      type="string"
                      size="small"

                      style={{ width: 200 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    checkpatient();
                    setSearch(true);

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
          </Grid>
        </Grid>


        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              {search ? (
                <div>
                  {  checkpatientid ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                          <TableCell align="center">Patient No.</TableCell>
                          <TableCell align="center">Name</TableCell>
                          <TableCell align="center">Age</TableCell>
                          <TableCell align="center">Disease</TableCell>
                          <TableCell align="center">MedicalCare</TableCell>
                          <TableCell align="center">Tel.</TableCell>
                          <TableCell align="center">BirthDay</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {patient.filter((filter: any) => filter.patientID.includes(patientid)).map((item: any) => (
                            <TableRow key={item.id}>
                               <TableCell align="center">{item.patientID}</TableCell>
                               <TableCell align="center">{item.name}</TableCell>
                               <TableCell align="center">{item.age}</TableCell>
                               <TableCell align="center">{item.edges?.disease?.name}</TableCell>
                               <TableCell align="center">{item.edges?.medicalcare?.name}</TableCell>
                               <TableCell align="center">{item.tel}</TableCell>
                               <TableCell align="center">{moment(item.birthday).format('DD/MM/YYYY')}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : patientid == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                              <TableCell align="center">Patient No.</TableCell>
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
                                  <TableCell align="center">{item.patientID}</TableCell>
                                  <TableCell align="center">{item.name}</TableCell>
                                  <TableCell align="center">{item.age}</TableCell>
                                  <TableCell align="center">{item.edges?.disease?.name}</TableCell>
                                  <TableCell align="center">{item.edges?.medicalcare?.name}</TableCell>
                                  <TableCell align="center">{item.tel}</TableCell>
                                  <TableCell align="center">{moment(item.birthday).format('DD/MM/YYYY')}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );

}