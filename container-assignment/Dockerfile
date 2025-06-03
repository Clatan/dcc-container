# --- Stage 1: Composer dependencies ---
FROM php:8.3-cli-alpine AS composer

# Install tools & composer
RUN apk add --no-cache git zip unzip curl && \
    curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# Set working dir & copy composer files
WORKDIR /app
COPY php/composer.json php/composer.lock ./

# Run composer install
RUN composer install --no-dev --prefer-dist --no-scripts --no-progress && \
    ls -la /app && ls -la /app/vendor

# --- Stage 2: Node dependencies & Vite build ---
FROM node:18-alpine AS frontend

WORKDIR /app
COPY php/package.json php/vite.config.js ./
RUN npm install

COPY php/resources ./resources
COPY php/public ./public
RUN npm run build

# --- Stage 3: Final PHP runtime ---
FROM php:8.3-fpm-alpine AS app

# Install PHP extensions and deps
RUN apk add --no-cache \
    bash \
    curl \
    git \
    zip \
    unzip \
    libpng-dev \
    libjpeg-turbo-dev \
    freetype-dev \
    oniguruma-dev \
    libxml2-dev \
    libwebp-dev \
    libxpm-dev \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install pdo pdo_mysql mbstring xml gd

# Set working directory
WORKDIR /var/www

COPY --from=composer /app/vendor /var/www/vendor

COPY php/app /var/www/app
COPY php/bootstrap /var/www/bootstrap
COPY php/config /var/www/config
COPY php/database /var/www/database
COPY php/public /var/www/public
COPY php/resources /var/www/resources
COPY php/routes /var/www/routes
COPY php/storage /var/www/storage
COPY php/tests /var/www/tests

# Copy important root-level files
COPY php/artisan /var/www/artisan
COPY php/composer.json /var/www/composer.json
COPY php/composer.lock /var/www/composer.lock
COPY php/phpunit.xml /var/www/phpunit.xml
COPY php/vite.config.js /var/www/vite.config.js
COPY php/package.json /var/www/package.json

# Copy built Vite assets
COPY --from=frontend /app/public/build ./public/build

# Set permissions (opsional, sesuai kebutuhan Laravel)
RUN chown -R www-data:www-data /var/www/storage /var/www/bootstrap/cache

EXPOSE 9000
CMD ["php-fpm"]