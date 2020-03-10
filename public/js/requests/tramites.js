import {showToastNotification} from "../alerts/notifications.js";

let table = $('#lista-tramites');

const ws = new WebSocket('ws://localhost:3000/ws');

ws.onopen = () => {
    console.log('conectado');
};

ws.onerror = () => {
    console.log('error');
};

ws.onmessage = e => {
    console.log(e.data);
    const msg = JSON.parse(e.data);
};

$('.btn-editar-tramite').on('click', function (e) {
    e.preventDefault();
    $.ajax({
        url: '/tramites/lista',
        method: 'GET',
        contentType: 'application/json',
        success: function (data) {
            ws.send(data);
        }
    });
});

$(document).ready(function () {
    $('#lista-tramites').DataTable({
        "processing": true,
        "ajax": {
            "url": "/tramites/lista",
            "contentType": "application/json",
            "type": "GET",
            "dataSrc": function (json) {
                return json;
            }
        },
        "columns": [
            { "data": "id" },
            { "data": "nombre" },
            { "data": "iniciales" },
            { "data": "tercera_edad", "render": function (data, type, row) {
                    if(row['tercera_edad'] === '1') {
                        return "SI";
                    }else {
                        return "NO";
                    }
                }},
            { "data": null, "render": function (data, type, row) {
                    return `
                        <div class="btn-group">
                            <a href="/tramites/editar/${row['id']}" class="btn btn-warning btn-editar-tramite">
                                <i class="fas fa-edit"></i>
                            </a>
                            <a href="/tramites/eliminar/${row['id']}" class="btn btn-danger btn-eliminar-tramite">
                                <i class="fas fa-trash-alt"></i>
                            </a>
                        </div>
                    `;
                } },
        ],
        "paging": true,
        "lengthChange": true,
        "searching": true,
        "ordering": true,
        "info": true,
        "autoWidth": true,
    });

    $('#iniciales').on('input', () => {
        let el = document.getElementById('iniciales');
        let p = el.selectionStart;
        el.value = el.value.toUpperCase();
        el.setSelectionRange(p, p);
    });

    $(document).on('click', '.btn-eliminar-tramite', function(e) {
        e.preventDefault();
        let deleteButton = $(this);
        let url = deleteButton.attr('href');
        let table = $('#lista-tramites').DataTable();

        Swal.fire({
            title: '¿Eliminar?',
            text: "¡No podrás revertir esto!",
            icon: 'question',
            showCancelButton: true,
            confirmButtonColor: '#d33',
            cancelButtonColor: '#3085d6',
            confirmButtonText: 'Si',
            cancelButtonText: 'No'
        }).then((result) => {
            if (result.value) {
                $.ajax({
                    url: url,
                    method: "DELETE",
                    contentType: 'application/json',
                    success: function () {
                        // table.row(deleteButton.parents('tr')).remove().draw();
                        table.ajax.reload();
                        showToastNotification('Eliminado correctamente', 'success');
                    },
                    error: function () {
                        showToastNotification('No se pudo eliminar', 'error');
                    }
                });
            }
        });
    });

    $('#form-nuevo-tramite').on('submit', e => {
        e.preventDefault();

        let url = $('#form-nuevo-tramite').attr('action');
        let table = $('#lista-tramites').DataTable();

        let nombre = $('input[name="nombre"]').val();
        let iniciales = $('input[name="iniciales"]').val();
        let terceraEdad = ($('input[name="tercera_edad"]').is(':checked')) ? '1' : '0';

        $.ajax({
            method: 'POST',
            url: url,
            data: JSON.stringify({nombre: nombre, iniciales: iniciales, tercera_edad: terceraEdad}),
            // dataType: 'json',
            contentType: 'application/json',
            beforeSend: function () {
                $('#btn-guardar-tramite').addClass('disabled');
                $('#btn-guardar-tramite > span.guardar').addClass('d-none');
                $('#btn-guardar-tramite > span.cargando').removeClass('d-none');
            },
            success: function () {
                document.getElementById('form-nuevo-tramite').reset();
                table.ajax.reload();
                showToastNotification('Trámite agregado correctamente', 'success');
                $('#btn-guardar-tramite').removeClass('disabled');
                $('#btn-guardar-tramite > span.guardar').removeClass('d-none');
                $('#btn-guardar-tramite > span.cargando').addClass('d-none');
            },
            error: function () {
                showToastNotification('No se pudo agregar el trámite', 'error');
                $('#btn-guardar-tramite').removeClass('disabled');
                $('#btn-guardar-tramite > span.guardar').removeClass('d-none');
                $('#btn-guardar-tramite > span.cargando').addClass('d-none');
            }
        });
    });
});