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
import { EntDentist } from '../../api/models/EntDentist';
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
  const [checkcardid, setcheckCardids] = useState(false);
  const [dentist, setDentist] = useState<EntDentist[]>([]);
  //--------------------------
  const [cardid, setCardids] = useState(String);
  const profile = { givenName: 'ระบบค้นหาข้อมูลทันตแพทย์' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getDentist = async () => {
      const res = await api.listDentist({ offset: 0 });
      setLoading(false);
      setDentist(res);
    };
    getDentist();
  }, [loading]);

  //-------------------
  const Cardidhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setcheckCardids(false);
    setCardids(event.target.value as string);

  };

  const cleardata = () => {
    setCardids("");
    setSearch(false);
    setcheckCardids(false);
    setSearch(false);

  }
  //---------------------
  const checkdentist = async () => {
    var check = false;
    dentist.map(item => {
      if (cardid != "") {
        if (item.cardid?.includes(cardid)) {
          setcheckCardids(true);
          alertMessage("success", "พบข้อมูล");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูล");
    }
    console.log(checkcardid)
    if (cardid == "") {
      alertMessage("info", "แสดงข้อมูลทันตแพทย์");
    }
  };

  return (

    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ค้นหาข้อมูลทันตแพทย์">
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
                <div style={{ background: "#0288D1", height: 50 }}>
                  <h1 style={
                    {
                      color: "#FFFFFF",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '30px',
                    }}>
                    ค้นหาข้อมูลทันตแพทย์
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}><strong>หมายเลขบัตรประชาน 13 หลักของทันตแพทย์ </strong></div>
                    <TextField
                      id="cardid"
                      value={cardid}
                      onChange={Cardidhandlehange}
                      type="string"
                      size="medium"

                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    checkdentist();
                    setSearch(true);

                  }}
                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#0288D1", height: 40 }}>
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
                  {  checkcardid ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                          <TableCell align="center">Dentist_ID</TableCell>
                          <TableCell align="center">Dentist_Name</TableCell>
                          <TableCell align="center">age</TableCell>
                          <TableCell align="center">BirthDay</TableCell>
                          <TableCell align="center">CartID</TableCell>
                          <TableCell align="center">Degree</TableCell>
                          <TableCell align="center">Gender</TableCell>
                          <TableCell align="center">Expert</TableCell>
                          <TableCell align="center">Experience</TableCell>
                          <TableCell align="center">Email</TableCell>
                          <TableCell align="center">Tel</TableCell>
                          <TableCell align="center">Password</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>
                          {dentist.filter((filter: any) => filter.cardid.includes(cardid)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.name}</TableCell>
                              <TableCell align="center">{item.age}</TableCell>
                              <TableCell align="center">{moment(item.birthday).format('DD/MM/YYYY HH:mm')}</TableCell>
                              <TableCell align="center">{item.cardid}</TableCell>
                              <TableCell align="center">{item.edges?.degree?.name}</TableCell>
                              <TableCell align="center">{item.edges?.gender?.name}</TableCell>
                              <TableCell align="center">{item.edges?.expert?.name}</TableCell>
                              <TableCell align="center">{item.experience}</TableCell>
                              <TableCell align="center">{item.email}</TableCell>
                              <TableCell align="center">{item.tel}</TableCell>
                              <TableCell align="center">{item.password}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : cardid == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                              <TableCell align="center">Dentist_ID</TableCell>
                          <TableCell align="center">Dentist_Name</TableCell>
                          <TableCell align="center">age</TableCell>
                          <TableCell align="center">BirthDay</TableCell>
                          <TableCell align="center">CartID</TableCell>
                          <TableCell align="center">Degree</TableCell>
                          <TableCell align="center">Gender</TableCell>
                          <TableCell align="center">Expert</TableCell>
                          <TableCell align="center">Experience</TableCell>
                          <TableCell align="center">Email</TableCell>
                          <TableCell align="center">Tel</TableCell>
                          <TableCell align="center">Password</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {dentist.map((item: any) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.name}</TableCell>
                              <TableCell align="center">{item.age}</TableCell>
                              <TableCell align="center">{moment(item.birthday).format('DD/MM/YYYY HH:mm')}</TableCell>
                              <TableCell align="center">{item.cardid}</TableCell>
                              <TableCell align="center">{item.edges?.degree?.name}</TableCell>
                              <TableCell align="center">{item.edges?.gender?.name}</TableCell>
                              <TableCell align="center">{item.edges?.expert?.name}</TableCell>
                              <TableCell align="center">{item.experience}</TableCell>
                              <TableCell align="center">{item.email}</TableCell>
                              <TableCell align="center">{item.tel}</TableCell>
                              <TableCell align="center">{item.password}</TableCell>
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