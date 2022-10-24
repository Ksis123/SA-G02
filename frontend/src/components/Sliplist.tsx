import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { SliplistInterface } from "../interfaces/ISliplist";
import { GetSlipList } from "../services/HttpClientServer";

function Sliplist() {
  const [slipList, setSliplist] = useState<SliplistInterface[]>([]);

  useEffect(() => {
    getSlipLists();
  }, []);

  const getSlipLists = async () => {
    let res = await GetSlipList();
    if (res) {
        setSliplist(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 60 },
    { field: "StudentList",headerName: "นักศึกษาที่ถูกคัดเลือก",width: 150, valueFormatter: (params) => params.value.ID,},
    { field: "Pay",headerName: "สถานะการโอนเงินทุนการศึกษา",width: 250,valueFormatter: (params) => params.value.Name,},
    { field: "Banking",headerName: "บัญชีธนาคาร",width: 150,valueFormatter: (params) => params.value.Commerce,},
    { field: "Total", headerName: "จำนวนเงิน", width: 100 ,},
    { field: "Slipdate", headerName: "วันที่และเวลา", width: 200 },
  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการบันทึกข้อมูลธุรกรรมการเงินทุนการศึกษา
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/sliplist/create"
              variant="contained"
              color="primary"
            >
              สร้างรายการธุรกรรมทุนการศึกษา
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={slipList}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
        <h4>Requirements</h4>
            <p>ระบบทุนการศึกษาเป็น ระบบสำหรับนักศึกษาผู้ประสงค์ขอรับทุนการศึกษาทางนักศึกษาที่ login เข้ามาเป็น User 
              ในระบบสามารถเลือกหาทุนการศึกษาที่ทางสนใจและคุณสมบัติตรงกับการรับทุนนั้น โดยมีรายละเอียดเงื่อนไขการรับทุนให้นักศึกษาได้ทราบซึ่งหากประสงค์
              สามารถส่งคำร้องขอรับทุนการศึกษา
            </p>  
            <p>
              นอกจากนี้ทาง ผู้ดูแลระบบที่ login เข้ามาเป็น Admin สามารถตรวจสอบบันทึกและแก้ไขรายการโอนเงินทุนการศึกษาของตัวนักศึกษาที่ถูกคัดเลือก
              ได้ผ่านทางระบบทุนการศึกษามหาวิทยาลัย โดยเมื่อขั้นตอนทำธุรกรรมการเงินผ่านทางธนาคารจนเงินทุนโอนเข้าสู่บัญชีนักศึกษาที่ถูกคัดเลือก
              แล้วทางระบบจะมีการบันทึกยอดการโอนเงิน เข้าสู่รายการโอนเงินทุนการศึกษาเพื่อตรวจสอบย้อนหลัง
            </p>
      </Container>
    </div>
  );
}

export default Sliplist;