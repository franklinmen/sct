import { showToastNotification} from "/public/js/alerts/notifications.js";

$('#btn-siguiente').on('click', e => {
    e.preventDefault();

    showToastNotification('Funciona Correctamente Siguiente', 'success');
});

$('#btn-llamar').on('click', e => {
    e.preventDefault();

    showToastNotification('Funciona Correctamente Llamada', 'success');
});