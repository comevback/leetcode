package Leetcode;

// leetcode-java/src/Leetcode/Sort_List_148.java
public class Sort_List_148 {
    public static void main(String[] args){
        // [4,2,1,3]
        ListNode head = new ListNode(4);
        ListNode node1 = new ListNode(2);
        ListNode node2 = new ListNode(1);
        ListNode node3 = new ListNode(3);
        head.next = node1;
        node1.next = node2;
        node2.next = node3;
        Sort_List_148 app = new Sort_List_148();
        ListNode res = app.sortList(head);
        while (res != null){
            System.out.print(res.val + " ");
            res = res.next;
        }
    }

    public ListNode sortList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode quick = new ListNode(0);
        ListNode slow = new ListNode(0);

        quick.next = head;
        slow.next = head;

        while (quick.next!=null){
            quick = quick.next;
            slow = slow.next;
            if (quick.next != null){
                quick = quick.next;
            }
        }

        ListNode right = slow.next;
        slow.next = null;

        ListNode resLeft = sortList(head);
        ListNode resRight = sortList(right);

        return merge(resLeft,resRight);
    }

    public ListNode merge(ListNode list1, ListNode list2){
        ListNode dummy = new ListNode(0);
        ListNode keep = dummy;

        while (list1 != null && list2 != null){
            if (list1.val<=list2.val){
                dummy.next = list1;
                dummy = dummy.next;
                list1 = list1.next;
            }else{
                dummy.next = list2;
                dummy = dummy.next;
                list2 = list2.next;
            }
        }

        if (list1 != null){
            dummy.next = list1;
        }

        if (list2 != null){
            dummy.next = list2;
        }

        return keep.next;
    }

    public static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
}
