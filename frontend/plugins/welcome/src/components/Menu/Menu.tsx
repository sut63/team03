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
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import ButtonGroup from '@material-ui/core/ButtonGroup';

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

const Homepage: FC<{}> = () => {
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
        <ButtonGroup size="small" aria-label="small outlined button group">
        <Button><Link color="textPrimary" href="/SavePatient" > บันทึกประวัติผู้ป่วย </Link></Button>
        <Button><Link color="textPrimary" href="/SaveDentist" > บันทึกข้อมูลทันตแพทย์ </Link></Button>
        <Button><Link color="textPrimary" href="/SaveMed" > บันทึกประวัติทันตกรรม </Link></Button>
        <Button><Link color="textPrimary" href="/SaveDenExpen" > บันทึกค่ารักษา </Link></Button>
        <Button><Link color="textPrimary" href="/SaveQueue" > บันทึกการจองคิว </Link></Button>
        </ButtonGroup>
        </center>

        <center>
        <ButtonGroup size="small" aria-label="small outlined button group">
        <Button><Link color="textPrimary" href="/SearchMed" > ค้นหาประวัติทันตกรรม </Link></Button>
        <Button><Link color="textPrimary" href="/SearchDenExpen" > ค้นหารายการค่ารักษา </Link></Button>
        </ButtonGroup>
        </center>

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
export default Homepage;
