import React, { useState } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import moment from 'moment';
import { Page, pageTheme, Header, Content} from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { EntAppointment } from '../../api';
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

  const [appointID, setAppointID] = useState(String);
  const [appointment, setAppointment] = useState<EntAppointment[]>([])
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);




  const AppointIDhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setAppointID(event.target.value as string);
  };

  const cleardata = () => {
    setAppointID("");
    setStatus(false);
    setAppointment([]);

  }

  

  const SearchAppointment = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchappointments?appointment=${appointID}`;
    const requestOptions = {
        method: 'GET',
    };
    fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
            console.log(data.data)
            setStatus(true);
            setAlert(false);
            setAppointment([]);

          if (data.data != null) {
                if(data.data.length >= 1) {
                  setStatus(true);
                  setAlert(true);
                  console.log(data.data)
                  setAppointment(data.data);
                }
            }
        });

}

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
                    <div className={classes.paper}>รหัสการนัดหมายผู้ป่วย ex.1</div>
                    <TextField
                      id="appointID"
                      value={appointID}
                      onChange={AppointIDhandlehange || ''}
                      type="string"
                      size="medium"

                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>

                <div></div>
                
                <Button
                  onClick={() => {
                    SearchAppointment();
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

            {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    แสดงข้อมูลการนัดหมาย
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
                <TableCell align="left">id</TableCell>
                <TableCell align="left">รหัสการนัดหมาย</TableCell>
                <TableCell align="left">รหัสผู้ป่วย</TableCell>
                <TableCell align="left">ชื่อผู้ป่วย</TableCell>
                <TableCell align="left">สาเหตุการนัดหมาย</TableCell>
                <TableCell align="left">วันที่และเวลา</TableCell>
                <TableCell align="left">Room</TableCell>
                <TableCell align="left">หมายเหตุ</TableCell>
              </TableRow>
            </TableHead>
          <TableBody>
            {appointment.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="left">{item.id}</TableCell>
                  <TableCell align="left">{item.AppointID}</TableCell>
                  <TableCell align="left">{item.edges?.Patient?.PatientID}</TableCell>
                  <TableCell align="left">{item.edges?.Patient?.Name}</TableCell>
                  <TableCell align="left">{item.Detail}</TableCell>
                  <TableCell align="left">{moment(item.Datetime).format('DD/MM/YYYY HH:mm')}</TableCell>
                  <TableCell align="left">{item.edges?.Room?.name}</TableCell>
                  <TableCell align="left">{item.Remark}</TableCell>
                                    
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