$(document).ready(function () {
    $('#body').toggleClass(localStorage.getItem('sc'));
    let cm = document.getElementById('collapse-menu');

    cm.addEventListener('click', (e) => {
        if($('#body').hasClass('sidebar-collapse')){
            localStorage.setItem('sc', '');
        } else {
            localStorage.setItem('sc', 'sidebar-collapse');
        }
    });
    $('#modulos').DataTable({
        "paging": true,
        "lengthChange": true,
        "searching": true,
        "ordering": true,
        "info": true,
        "autoWidth": true,
    });
    // $('#lista-tramites').DataTable({
    //     "paging": true,
    //     "lengthChange": true,
    //     "searching": true,
    //     "ordering": true,
    //     "info": true,
    //     "autoWidth": true,
    // });
});