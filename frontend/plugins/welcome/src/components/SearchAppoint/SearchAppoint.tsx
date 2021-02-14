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
import Swal from 'sweetalert2'
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { EntAppointment } from '../../api';


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
  position: 'center',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,

});


export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  //---------------------------
  const [checkAppointID, setCheckAppointID] = useState(false);
  const [appointment, setAppointment] = useState<EntAppointment[]>([]);
  const [appointID, setAppointID] = useState(String);

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getAppointment = async () => {
      const res = await api.listAppointment({ offset: 0 });
      setLoading(false);
      setAppointment(res);
    };
    getAppointment();
  }, [loading]);

  //-------------------
  const AppointIDhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setCheckAppointID(false);
    setAppointID(event.target.value as string);

  };

  const cleardata = () => {
    setAppointID("");
    setSearch(false);
    setCheckAppointID(false);
    setSearch(false);

  }
  //---------------------
  const CheckAppointment = async () => {
    var check = false;
    appointment.map(item => {
      if (appointID != "") {
        if (item.appointID?.includes(appointID)) {
          setCheckAppointID(true);
          alertMessage("success", "พบข้อมูล");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูล");
    }
    console.log(checkAppointID)
    if (appointID == "") {
      alertMessage("info", "แสดงข้อมูลการนัดหมายผู้ป่วย");
    }
  };

  return (

    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ค้นหาการนัดหมายผู้ป่วย">
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
                <div style={{ background: "#3792cb", height: 50 }}>
                  <h1 style={
                    {
                      color: "#FFFFFF",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '30px',
                    }}>
                    ค้นหาการนัดหมายผู้ป่วย
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>รหัสการนัดหมายผู้ป่วย ex.A00001</div>
                    <TextField
                      id="appointID"
                      value={appointID}
                      onChange={AppointIDhandlehange}
                      type="string"
                      size="medium"

                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>

                <div></div>
                
                <Button
                  onClick={() => {
                    CheckAppointment();
                    setSearch(true);

                  }}
                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#3792cb", height: 40 }}>
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
                  style={{ background: "#3792cb", height: 40 }}>
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
          </Grid>
        </Grid>


        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              {search ? (
                <div>
                  {  checkAppointID ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                          <TableCell align="left">No.</TableCell>
                          <TableCell align="left">Appointment ID</TableCell>
                          <TableCell align="left">Patient</TableCell>
                          <TableCell align="left">Detail</TableCell>
                          <TableCell align="left">Date-Time</TableCell>
                          <TableCell align="left">Room</TableCell>
                          <TableCell align="left">Remark</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {appointment.filter((filter: any) => filter.appointID.includes(appointID)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="left">{item.id}</TableCell>
                              <TableCell align="left">{item.appointID}</TableCell>
                              <TableCell align="left">{item.edges?.patient?.name}</TableCell>
                              <TableCell align="left">{item.detail}</TableCell>
                              <TableCell align="left">{moment(item.datetime).format('DD/MM/YYYY HH:mm')}</TableCell>
                              <TableCell align="left">{item.edges.room.name}</TableCell>
                              <TableCell align="left">{item.remark}</TableCell>
                              
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : appointID == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                          <TableHead>
                            <TableRow>
                            <TableCell align="left">No.</TableCell>
                            <TableCell align="left">Appointment ID</TableCell>
                            <TableCell align="left">Patient</TableCell>
                            <TableCell align="left">Detail</TableCell>
                            <TableCell align="left">Date-Time</TableCell>
                            <TableCell align="left">Room</TableCell>
                            <TableCell align="left">Remark</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>

                        {appointment.map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="left">{item.id}</TableCell>
                              <TableCell align="left">{item.appointID}</TableCell>
                              <TableCell align="left">{item.edges?.patient?.name}</TableCell>
                              <TableCell align="left">{item.detail}</TableCell>
                              <TableCell align="left">{moment(item.datetime).format('DD/MM/YYYY HH:mm')}</TableCell>
                              <TableCell align="left">{item.edges.room.name}</TableCell>
                              <TableCell align="left">{item.remark}</TableCell>
                              
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