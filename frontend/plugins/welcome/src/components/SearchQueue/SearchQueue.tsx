import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis'
import Swal from 'sweetalert2'
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { EntQueue } from '../../api/models/EntQueue';

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
  const [checkqueueID, setcheckQueueIDs] = useState(false);
  const [queue, setQueue] = useState<EntQueue[]>([]);
  //--------------------------
  const [queueID, setQueueIDs] = useState(String);
  const profile = { givenName: 'ระบบค้นหาใบจองคิวผู้ป่วย' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getQueue = async () => {
      const res = await api.listQueue({ offset: 0 });
      setLoading(false);
      setQueue(res);
    };
    getQueue();
  }, [loading]);

  //-------------------
  const QueueIDhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setcheckQueueIDs(false);
    setQueueIDs(event.target.value as string);

  };

  const cleardata = () => {
    setQueueIDs("");
    setSearch(false);
    setcheckQueueIDs(false);
    setSearch(false);

  }
  //---------------------
  const checkqueue = async () => {
    var check = false;
    queue.map(item => {
      if (queueID != "") {
        if (item.queueID?.includes(queueID)) {
          setcheckQueueIDs(true);
          alertMessage("success", "พบข้อมูล");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูล");
    }
    console.log(checkqueue)
    if (queueID == "") {
      alertMessage("info", "แสดงข้อมูลประวัติทันตกรรม");
    }
  };

  return (

    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ค้นหาใบจองคิวผู้ป่วย">
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
                <div style={{ background: "#37d67a", height: 60 }}>
                  <h1 style={
                    {
                      color: "#FFFFFF",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '30px',
                    }}>
                    ค้นหาใบจองคิวผู้ป่วย
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}><strong>หมายเลขใบจองคิว ex.Q001</strong></div>
                    <TextField
                      id="queueID"
                      value={queueID}
                      onChange={QueueIDhandlehange}
                      type="string"
                      size="small"

                      style={{ width: 200 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    checkqueue();
                    setSearch(true);

                  }}
                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#37d67a", height: 40 }}>
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
                  {  checkqueueID ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                          <TableCell align="center">No.</TableCell>
                          <TableCell align="center">Queue_ID</TableCell>
                          <TableCell align="center">Patient</TableCell>
                          <TableCell align="center">Phone</TableCell>
                          <TableCell align="center">Dental</TableCell>
                          <TableCell align="center">Dentist</TableCell>
                          <TableCell align="center">Date</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {queue.filter((filter: any) => filter.queueID.includes(queueID)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.queueID}</TableCell>
                              <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                              <TableCell align="center">{item.phone}</TableCell>
                              <TableCell align="center">{item.dental}</TableCell>
                              <TableCell align="center">{item.edges?.dentist?.name}</TableCell>
                              <TableCell align="center">{moment(item.queueTime).format('DD/MM/YYYY HH:mm')}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : queueID == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                                <TableCell align="center">No.</TableCell>
                                <TableCell align="center">Queue_ID</TableCell>
                                <TableCell align="center">Patient</TableCell>
                                <TableCell align="center">Phone</TableCell>
                                <TableCell align="center">Dental</TableCell>
                                <TableCell align="center">Dentist</TableCell>
                                <TableCell align="center">Date</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {queue.map((item: any) => (
                                <TableRow key={item.id}>
                                  <TableCell align="center">{item.id}</TableCell>
                                  <TableCell align="center">{item.queueID}</TableCell>
                                  <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                                  <TableCell align="center">{item.phone}</TableCell>
                                  <TableCell align="center">{item.dental}</TableCell>
                                  <TableCell align="center">{item.edges?.dentist?.name}</TableCell>
                                  <TableCell align="center">{moment(item.queueTime).format('DD/MM/YYYY HH:mm')}</TableCell>
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