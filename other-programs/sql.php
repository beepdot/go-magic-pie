<!DOCTYPE html>
<html>
<head>
<style>
table {
    width: 100%;
    border-collapse: collapse;
}

table, td, th {
    border: 1px solid black;
    padding: 5px;
}

th {text-align: left;}
</style>
</head>
<body>

<?php

$con = mysqli_connect('localhost','root','root','devcon');
if (!$con) {
    die('Could not connect: ' . mysqli_error($con));
}

mysqli_select_db($con,"devcon");
$sql="SELECT * FROM userdata WHERE id = 4";
$result = mysqli_query($con,$sql);

echo "<table>
<tr>
<th>Image</th>
<th>Name</th>
<th>VisitorID</th>
<th>Stall</th>
</tr>";
while($row = mysqli_fetch_array($result)) {
    echo "<tr>";
    echo "<td>" . $row['image'] . "</td>";
    echo "<td>" . $row['name'] . "</td>";
    echo "<td>" . $row['visitorid'] . "</td>";
    echo "<td>" . $row['stalls'] . "</td>";
    echo "</tr>";
}
echo "</table>";
mysqli_close($con);
?>
</body>
</html> 
