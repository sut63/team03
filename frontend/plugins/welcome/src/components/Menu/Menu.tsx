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
import MenuList from '@material-ui/core/MenuList';
import MenuItem from '@material-ui/core/MenuItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import Grid, { GridSpacing } from '@material-ui/core/Grid';
import SearchIcon from '@material-ui/icons/Search';
import SaveIcon from '@material-ui/icons/Save';
const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    justifyContent: 'center'
  },
  paper: {
    height: 530,
    width: 800,
    backgroundImage: 'url(https://www.nicepng.com/png/detail/38-387168_kids-transparent-dentist-clip-art-free-dentist-png.png)',
  },
  control: {
    padding: theme.spacing(2),
  },
  posi: {
      width: '100%',
      height: '6%',
      justifyContent: 'left',
      alignItems: 'center'
  },
  menu_root: {
    width: 300,
  },

})
);

const Homepage: FC<{}> = () => {
  const classes = useStyles();
  const [spacing] = React.useState<GridSpacing>(2);
  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ระบบทันตกรรม">
        <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
           </Button>
      </Header>
      <Content>

      <ContentHeader title="Menu (เมนู)"> </ContentHeader>
        
      <Grid container justify="space-around" alignItems="stretch" spacing={spacing}>
          
            <MenuList className={classes.menu_root}>
              <MenuItem>
                  <ListItemIcon>
                    <SaveIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SavePatient" >บันทึกประวัติผู้ป่วย</Link></MenuItem>

              <MenuItem>
                  <ListItemIcon>
                    <SaveIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SaveDentist" >บันทึกข้อมูลทันตแพทย์</Link></MenuItem>
              
              <MenuItem>
                  <ListItemIcon>
                    <SaveIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SaveMed" >บันทึกประวัติทันตกรรม</Link></MenuItem>

              <MenuItem>
                  <ListItemIcon>
                    <SaveIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SaveDenExpen" >บันทึกค่ารักษา</Link></MenuItem>

              <MenuItem>
                  <ListItemIcon>
                    <SaveIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SaveQueue" >บันทึกการจองคิว</Link></MenuItem>


              <MenuItem>
                  <ListItemIcon>
                    <SearchIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/" >ค้นหาประวัติผู้ป่วย</Link></MenuItem>

              <MenuItem>
                  <ListItemIcon>
                    <SearchIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SearchMed" >ค้นหาประวัติทันตกรรม</Link></MenuItem>

              <MenuItem>
                  <ListItemIcon>
                    <SearchIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SavePatient" >ค้นหาการข้อมูลทันตแพทย์</Link></MenuItem>          

              <MenuItem>
                  <ListItemIcon>
                    <SearchIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SearchDenExpen" >ค้นหารายการค่ารักษา</Link></MenuItem>

              <MenuItem>
                  <ListItemIcon>
                    <SearchIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/" >ค้นหาการจองคิว</Link></MenuItem>

              <MenuItem>
                  <ListItemIcon>
                    <SearchIcon color="primary" fontSize="small" />
                  </ListItemIcon>
                    <Link color="textPrimary" href="/SearchAppoint" >ค้นหาการนัดหมายผู้ป่วย</Link></MenuItem>
            </MenuList>
         
          {[0].map((value) => (
            <Grid key={value} item>
              <Paper className={classes.paper} />
            </Grid>
          ))}
      </Grid>
    
      </Content>
    </Page>
  );
};
export default Homepage;
