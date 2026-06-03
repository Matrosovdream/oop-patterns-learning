<?php

use App\Http\Controllers\OrderingController;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});

// The Ordering bounded context, exercised over HTTP. Run `docker compose up laravel`
// and open http://localhost:8000/ordering — you'll get the CQRS read model as JSON.
// See php-track/learning-plan lessons 12-20 for what's happening behind this route.
Route::get('/ordering', [OrderingController::class, 'demo']);
