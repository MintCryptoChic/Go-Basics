package main

import "fmt"

type Employee struct{
    Id string
    Name string
    Salary int
}

type Project struct{
    Project_Id string
    Date_Started string
    Assigned_Employees []string 
}

type Company struct{
    Company_Name string
    Year_Established int
    Employees []Employee
    Projects []Project
}

func (comp Company) CompInfo() string{
    fmt.Println(comp.Company_Name, comp.Year_Established)
    for _, emp:= range comp.Employees{
        fmt.Println(emp.Name, emp.Salary)
    }
    for _, proj:= range comp.Projects{
        fmt.Println(proj.Date_Started)
    }
    return "__"
}

func main(){
    comp:= Company{
        Company_Name: "AYS Software Solutions",
        Year_Established: 2019,
        Employee:[]Employees {
            Employees{
                "Id": "001" ,
                "Name": "Pragya",
                "Salary": 80000,
            },
    
            Employees{
                "Id": "002",
                "Name": "Pratyaksha",
                "Salary": 85000,
            },
    
            Employees{
                "Id":"003",
                "Name": "Sharwari",
                "Salary": 70000,
            },
        },
    
        Project:[]Projects {
            Projects{
                "Project_Id": "101",
                "Date_Started":"Dec 15",
                "Assigned_Employees":[]string{"002","003"},
            },
    
            Projects{
                "Project_Id": "102",
                "Date_Started": "Oct 31",
                "Assigned_Employees":[]string{"003"},
            },
    
            Projects{
                "Project_Id":"103",
                "Date_Started":"Aug 16",
                "Assigned_Employees":[]string{"001", "002"},
            },
        }
    }

    fmt.Print(comp.CompInfo())
}