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
import { EntMedicalfile } from '../../api/models/EntMedicalfile';
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
  const [checkmedno, setcheckMednos] = useState(false);
  const [medicalfile, setMedicalfile] = useState<EntMedicalfile[]>([]);
  //--------------------------
  const [medno, setMednos] = useState(String);
  const profile = { givenName: 'ระบบค้นหาประวัติทันตกรรม' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getMedicalfile = async () => {
      const res = await api.listMedicalfile({ offset: 0 });
      setLoading(false);
      setMedicalfile(res);
    };
    getMedicalfile();
  }, [loading]);

  //-------------------
  const Mednohandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setcheckMednos(false);
    setMednos(event.target.value as string);

  };

  const cleardata = () => {
    setMednos("");
    setSearch(false);
    setcheckMednos(false);
    setSearch(false);

  }
  //---------------------
  const checkmedicalfile = async () => {
    var check = false;
    medicalfile.map(item => {
      if (medno != "") {
        if (item.medno?.includes(medno)) {
          setcheckMednos(true);
          alertMessage("success", "พบข้อมูล");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูล");
    }
    console.log(checkmedno)
    if (medno == "") {
      alertMessage("info", "แสดงข้อมูลประวัติทันตกรรม");
    }
  };

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
                <div style={{ background: "#3366CC", height: 60 }}>
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
                    <div className={classes.paper}><strong>หมายเลขทันตกรรม ex.M001</strong></div>
                    <TextField
                      id="medno"
                      value={medno}
                      onChange={Mednohandlehange}
                      type="string"
                      size="small"

                      style={{ width: 200 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    checkmedicalfile();
                    setSearch(true);

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
                  style={{ background: "#FFFFE0", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#000000",
                        padding: '0 25px',

                      }
                    }>
                    เคลียร์ข้อมูล
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
                  {  checkmedno ? (
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

                          {medicalfile.filter((filter: any) => filter.medno.includes(medno)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.medno}</TableCell>
                              <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                              <TableCell align="center">{item.detial}</TableCell>
                              <TableCell align="center">{item.edges?.dentist?.name}</TableCell>
                              <TableCell align="center">{moment(item.addedTime).format('DD/MM/YYYY HH:mm')}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : medno == "" ? (
                      <div>
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
                                  <TableCell align="center">{item.medno}</TableCell>
                                  <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                                  <TableCell align="center">{item.detial}</TableCell>
                                  <TableCell align="center">{item.edges?.dentist?.name}</TableCell>
                                  <TableCell align="center">{moment(item.addedTime).format('DD/MM/YYYY HH:mm')}</TableCell>
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