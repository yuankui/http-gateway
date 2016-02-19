// Place all the behaviors and hooks related to the matching controller here.
// All this logic will automatically be available in application


$(document).ready(function () {
    $('table').DataTable({
        "paging": false,
        "searching": true,
        "info": false,
        "lengthChange": false,
        "language": {
            "lengthMenu": "x _MENU_ records"
        },
        "columnDefs": [
            { "searchable": false, "targets": 0 },
            { "searchable": false, "targets": 2 }
        ]
    });
});

$(function() {
    table = $('table').dataTable();
    $('#searchBox').keyup(function(){
        var query = $(this).val();
        table.fnFilter(query);
    });
})

