<?php
/*
 * Plugin Name:       Simple Barcode Generator
 * Plugin URI:        https://krakenhub.net/
 * Description:       Create basic barcode to be printed taking the SKU field of woocommerce
 * Version:           0.0.1
 * Author:            KRAKENHUB
 * Author URI:        https://krakenhub.net
 * License:           GPL v2 or later
 * License URI:       https://www.gnu.org/licenses/gpl-2.0.html
 */
/**
 * Activate the plugin.
 */
function kh_barcode_generator_activate() { 
    
}
register_activation_hook( __FILE__, 'kh_barcode_generator_activate' );

/**
 * Deactivation hook.
 */
function kh_barcode_generator_deactivate() {

}
register_deactivation_hook( __FILE__, 'kh_barcode_generator_deactivate' );

// Add product new column in administration
function kh_add_print_barcode_column( $columns ) {
    //add column
    $columns['print_barcode'] = __( 'Print Barcode', 'woocommerce' );

    return $columns;
}
add_filter( 'manage_edit-product_columns', 'kh_add_print_barcode_column', 10, 1 );

function kh_add_print_barcode_column_content( $column, $postid ) {
    if ( $column == 'print_barcode' ) {
        // Get product object
        $product = wc_get_product( $postid );
		
        // Get Product Variations
        $sku = $product->get_sku();
        echo '<center><img width="80px" src="https://barcode.khat.es/api/generate?v='.$sku.'"><br><a href="#" onclick="printBarcode('.$sku.');return false;">Print Barcode</a><br></center>';
    }
}
add_action( 'manage_product_posts_custom_column', 'kh_add_print_barcode_column_content', 10, 2 );

function kh_enqueue_admin_files($hook) {
    // Only add to the edit.php admin page.
    // See WP docs.
    if ('edit.php' !== $hook) {
        return;
    }
    wp_enqueue_script('kh_js_script', plugin_dir_url(__FILE__) . '/kh-functions.js');
}

add_action('admin_enqueue_scripts', 'kh_enqueue_admin_files');