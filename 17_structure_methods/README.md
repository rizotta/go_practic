**Задание:**
- Предположим, что теперь нам необходимо не позволить внешнему коду предоставлять нам реализацию свойства CalcDiscount, а жестко задать реализацию в привязке к структуре Customer, для этого сделайте CalcDiscount методом, а не свойством структуры Customer:
    - Логику оставить, как была в функции CalcDiscount.
    - Константу DEFAULT_DISCOUNT перенести в пакет internal.
    - Свойство Discount сделать нередактируемым вне пакета internal.