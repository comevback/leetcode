package Leetcode;

public class Swap_Nodes_in_Pairs_24 {
    public static void main(String[] args) {
        // [1,2,3,4]
        ListNode head = new ListNode(1);
        ListNode node1 = new ListNode(2);
        ListNode node2 = new ListNode(3);
        ListNode node3 = new ListNode(4);
        head.next = node1;
        node1.next = node2;
        node2.next = node3;

        Swap_Nodes_in_Pairs_24 app = new Swap_Nodes_in_Pairs_24();
        ListNode res = app.swapPairs(head);
        while (res != null) {
            System.out.print(res.val + " ");
            res = res.next;
        }
    }
    public ListNode swapPairs(ListNode head) {
        if (head == null){
            return head;
        }
        if (head.next == null){
            return head;
        }

        ListNode dummy = new ListNode(0,head.next);
        ListNode before = new ListNode(0,head);
        ListNode left = head;
        ListNode right = head.next;

        while(left != null && right != null){
            before.next = right;
            left.next = right.next;
            right.next = left;
            before = left;
            left = left.next;
            if (left != null){
                right = left.next;
            }
        }

        return dummy.next;
    }

    public static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
}

