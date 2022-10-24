import React, { useState, useEffect } from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import CssBaseline from '@mui/material/CssBaseline';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import { styled } from '@mui/material/styles';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { DataGrid, GridColDef, GridValueGetterParams } from '@mui/x-data-grid';
import AdbIcon from '@mui/icons-material/Adb';
import MenuItem from '@mui/material/MenuItem';
import Tooltip from '@mui/material/Tooltip';
import Avatar from '@mui/material/Avatar';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TablePagination from '@mui/material/TablePagination';
import { getValue } from '@mui/system';
import { Label } from '@mui/icons-material';
import FormControl from "@mui/material/FormControl";
import Select, { SelectChangeEvent } from '@mui/material/Select';
import InputLabel from '@mui/material/InputLabel';
import ReceiptIcon from '@mui/icons-material/Receipt';

import { ScholarshipStatusInterface } from '../interfaces/IScholarshipstatus';
import { ScholarshipTypesInterface } from '../interfaces/IScholarshiptypes';
import { ScholarshipInterface } from '../interfaces/IScholarship';

import {
    GetScholarship,
    GetScholarshipStatus,
    GetTypes,
    Scholarship,
  
} from "../services/HttpClientServer";
// function createData(
//   id: number,
//   Scholarship_Name: string,
//   Type: number,
//   status: number,
//   details: string,
// ) {
//   return { id,Scholarship_Name, Type, status, details };
// }

function createData(
  name: string,
  type: string,
  status: string,
  details: string,

)  {
  return {name,type,status,details};
}

function ScholarshipCreate() {

  const [rows, setRows] = useState<{
    name: string,
    type: string,
    status: string,
    details: string,
  }[]>([])

  const [ref, setRef] = React.useState(false);

  // const rows = [
  //   createData("1","2","3","4"),
  //  ];
  //  rows.push(createData(Scholarship_Name,types,status,scholarship_details))
  const [scholarshipstatuses,  setScholarshipstatus] = useState<ScholarshipStatusInterface[]>([]);
  const [Scholarship_Name, setScholarship_Name] = useState<string>("");
  const [scholarshipstypees, setScholarshipstype] = useState<ScholarshipTypesInterface[]>([]);
  const [scholarship_details, setScholarship_Details] = useState<string>("");
  const [status, setStatus] = React.useState('');
  const [types, setTypes] = React.useState('');
  //const [rowes, setRows] = useState('');
  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof scholarships;
    setScholarships({
      ...scholarships,
      [name]: event.target.value,
    });
  };
 
  const handleChangeStatus = (event: SelectChangeEvent) => {
    setStatus(event.target.value as string);
  };

  const handleChangeTypes = (event: SelectChangeEvent) => {
    setTypes(event.target.value as string);
  };

  const [scholarships, setScholarships] = useState<ScholarshipInterface>({
    Name: "", Details: ""
  });

  const update = () => {
    const temp = rows
    temp.push(createData(Scholarship_Name,types.toString(),status.toString(),scholarship_details))
    console.log(temp)
    setRows(temp)
    setRef(!ref)
  console.log(Scholarship_Name,types,status,scholarship_details);
  };

  useEffect(()=>{
    console.log(rows)
  },[rows,ref])

  const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(2),
    marginTop: 5,
    textAlign: 'center',
    color: theme.palette.text.secondary,
  }));

  const getScholarshipStatus = async () => {
    let res = await GetScholarshipStatus();
    scholarships.StatusID = res.ID;
    if (res) {
      setScholarshipstatus(res);
    }
  };

  const getScholarshipsTypes = async () => {
    let res = await GetTypes();
    scholarships.TypeID = res.ID;
    if (res) {
      setScholarshipstatus(res);
    }
  };
  const  getScholarship = async () => {
    let res = await GetScholarship();
    if (res) {
      setScholarships(res);
    } 
  };

  useEffect(() => {
      getScholarshipStatus();
      getScholarshipsTypes();
      getScholarship();
    }, []);
  return (
    <div>
    <Box sx={{ flexGrow: 1 }}>
  </Box>
  <Container maxWidth = "lg">
        {/* <Box sx={{ bgcolor: '#cfe8fc', height: '100vh', width: '190vh' }} /> */}
        <Paper >
        <Box
        display = {"flex"}
      sx={{
          marginTop: 2,
          paddingX:  2,
          paddingY:  4,
      }}
      >
     <h2>Manage Scholarship</h2>    
    </Box>
    <hr />
    <Grid container spacing={0}>
        <Grid xs={3}>
        <TextField fullWidth id="Scholarship_Name" type = "string" label="Scholarship Name" variant="outlined"  
        onChange ={(event) => setScholarship_Name(event.target.value)}
        />
          {/* <Item>xs=3</Item> */}
        </Grid>
        <Grid xs={3}>
        <FormControl fullWidth>
        <InputLabel id="demo-simple-select-label">Types</InputLabel>
        <Select
          labelId="เลือกประเภททุน"
          id="demo-simple-select"
          value={types}
          label="Status"
          onChange={handleChangeTypes}
        >
          <MenuItem value={1} >ทุนให้เปล่า</MenuItem>
          <MenuItem value={2}>ทุนต่อเนื่อง</MenuItem>
          <MenuItem value={3}>ทุนต่างประเทศ</MenuItem>
        </Select>
      </FormControl>

       
        {/* <FormControl fullWidth variant="outlined">
        <Select
                native
                value={scholarships.TypeID +""}
                onChange={handleChange}
                inputProps={{
                  name: "TypesID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภททุน
                </option>
                {scholarshipstypees.map((item:TypesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Type}
                  
                  
                  </option>
                ))}
              </Select>
            </FormControl>
          <Item>xs=3</Item> */}
        </Grid>
        <Grid xs={3}>
        <FormControl fullWidth>
        <InputLabel id="demo-simple-select-label">Status</InputLabel>
        <Select
          labelId="เลือกสถานะ"
          id="demo-simple-select"
          value={status}
          label="Status"
          onChange={handleChangeStatus}
        >
          <MenuItem value={1} key = {"เปิดรับอยู่"}>เปิดรับอยู่</MenuItem>
          <MenuItem value={2}>ปิดแล้วค้าบ</MenuItem>
          <MenuItem value={3}>อัพเดตอยู่นะ</MenuItem>
        </Select>
      </FormControl>

        {/* <FormControl fullWidth variant="outlined">
        <Select
                native
                value={scholarships.StatusID +""}
                onChange={handleChange}
              
                inputProps={{
                  name: "StatusID",
                }}
              >
            
                <option aria-label="None" value="">
                  กรุณาเลือกสถานะ
                </option>
                {scholarshipstatuses.map((item:StatusInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Status}
                  </option>
                ))}
        
              </Select>
            </FormControl> */}

        </Grid>
        <Grid xs={3}>
        <Button sx={{paddingY: 1.7,}} fullWidth variant="outlined" size="large"
         onClick={update}
        >
          Updates
        </Button>
          {/* <Item>xs=3</Item> */}
        </Grid>
        <Grid xs={12} sx ={{height: 200,}}>
        <TextField
          fullWidth
          id="details"
          type = "string"
          onChange ={(event) => setScholarship_Details(event.target.value)}
          label="Details"
          multiline
          rows={7}
          variant="outlined"
        
        />
        {/* <TextField  fullWidth  id="details" label="Details" multiline maxRows={7} variant="outlined" /> */}
          {/* <Item>xs=12</Item> */}
        </Grid>
        <Grid xs={12}>
        
          {/* <Item>xs=12</Item> */}
        </Grid>
    
        <Grid xs={9}>
       
        </Grid>
        <Grid xs={3}>
        <Button sx={{paddingY: 1.7,}} fullWidth  variant="outlined" size="large" onClick={update}>
          บันทึกข้อมูล
        </Button>
        </Grid>
      </Grid>
      
      </Paper >
     
      <Grid xs = {12}>
     
    </Grid>
   
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>รายชื่อทุน</TableCell>
            <TableCell align="left">ประเภท</TableCell>
            <TableCell align="left">สถานะ</TableCell>
            <TableCell align="left">รายละเอียด</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow
              key={row.name}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                {row.name}
              </TableCell>
              <TableCell align="left">{row.type}</TableCell>
              <TableCell align="left">{row.status}</TableCell>
              <TableCell align="left">{row.details}</TableCell>
           
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>

      </Container>
    
 </div>
   
  );
                }
 




export default ScholarshipCreate;