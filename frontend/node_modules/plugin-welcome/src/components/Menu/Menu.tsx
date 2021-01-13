import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid, { GridSpacing } from '@material-ui/core/Grid';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import { Link as RouterLink } from 'react-router-dom';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    justifyContent: 'center'
  },
  paper: {
    height: 440,
    width: 1250,
    backgroundImage: 'url(https://image.freepik.com/free-vector/healthy-kawaii-tooth-with-dental-braces-cute-cartoon-character_74565-467.jpg)',
  },
  control: {
    padding: theme.spacing(2),
  },
 
})
);

const Menu: FC<{}> = () => {
  const classes = useStyles();

  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ระบบทันตกรรม">
        <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
           </Button>
      </Header>
      <Content>
        
        <center>
        <ContentHeader title="Menu"> </ContentHeader>
        <ButtonGroup variant="contained" color="primary" aria-label="contained primary button group">
        <Link component={RouterLink} to="/SavePatient">
        <Button>บันทึกประวัติผู้ป่วย</Button></Link>
        <Link component={RouterLink} to="/SaveDentist">
        <Button>บันทึกข้อมูลแพทย์</Button></Link>
        <Link component={RouterLink} to="/SaveMed">
        <Button>บันทึกประวัติทันตกรรม</Button></Link>
        <Link component={RouterLink} to="/SaveDenExpen">
        <Button>บันทึกใบเสร็จค่ารักษา</Button></Link>
        <Link component={RouterLink} to="/SaveQueue">
        <Button>บันทึกการจองคิว</Button></Link>
        </ButtonGroup>
        </center>

        <Breadcrumbs aria-label="breadcrumb" >
        </Breadcrumbs>
        <br></br>
        <Grid container className={classes.root} spacing={2}>
      <Grid item xs={12}>
        <Grid container justify="center" >
          {[0].map((value) => (
            <Grid key={value} item>
              <Paper className={classes.paper} />
            </Grid>
          ))}
        </Grid>
      </Grid>
      <Grid item xs={12}>
      </Grid>
    </Grid>
      </Content>
    </Page>
  );
};

export default Menu;