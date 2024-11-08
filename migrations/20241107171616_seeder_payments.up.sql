INSERT INTO payment_sources (is_active,
                             code,
                             description)
VALUES (TRUE, 'MIDTRANS', 'Payment Gateway Midtrans');

INSERT INTO payment_methods (is_active,
                             payment_fee,
                             payment_fee_type,
                             code,
                             category,
                             name,
                             image)
VALUES (TRUE, 2.0, 'percentage', 'GP', 'E-Wallet', 'GoPay', 'gopay.png'),
       (TRUE, 2.0, 'percentage', 'SP', 'E-Wallet', 'ShopeePay', 'gopay.png'),
       (TRUE, 0.7, 'percentage', 'QR', 'QRIS', 'Qris', 'qris.png');