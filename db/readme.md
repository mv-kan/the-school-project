# Database structure

## Break down the problem

We break down **the problem** (in main readme) to simple facts like this:

Kan study at Awans
Kan is in TiB2 class
Kan's faculty is IT technology
Kan lives in Assystent 
Assystent fees are 385 (two persons) or 398 (three persons)
Kan pays evert month assystent's fee in amount 385 zl
Payments due very month

<Pupil> study at Awans
<Pupil> is in <Class>
<Pupil>'s facylty is <Faculty>
<Pupil> lives in <Dorm>
<Dorm>'s fees are <Price> 
<Pupil> make every month <Payment>
<Payment> due every month
<Supervisor> is assigned to every <pupil>
Payments due <Date>

if <pupil> has 18 y/o supervisor or has 
other supervisor not from awans then this <pupil>
does not have <supervisor>

In <available room> can not be more than <number> of pupils

if <pupil> has <supervisor> and <supervisor> is provided by Awans then <pupil> makes <payment> to pay for <supervisor> 
Each <Dormitory> has limited number of <available room>s

So each placeholder may be entity type (<Pupil>) and may not (<Date>).

## ER model and Relational model
After breaking down **the problem** we can create er model and relational model:

### ER model
![ER model](./diagrams/TheSchoolDiagram-ER%20model.jpg)

### Relational model
![Relational model](./diagrams/TheSchoolDiagram-Relational%20Model.jpg)

Original diagrams are [here](https://drive.google.com/file/d/1sxZLKAwHTJC1BVuSwDg3Cgny6A4BFn4V/view?usp=sharing)

## Migration 

As per usual [the tool](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
and in makefile in root folder there are commands to run migrations